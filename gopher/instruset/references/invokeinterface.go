package references

import (
	"gopher/instruset/base"
	"gopher/rtdata"
	"gopher/rtdata/heap"
)

type INVOKE_INTERFACE struct {
	index uint
	// count uint8  // parameters slots
	// zero uint8   // special for some oracle jvm
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count of slot
	reader.ReadUint8() // must be zero
}

func (self *INVOKE_INTERFACE) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointException")
	}

	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
