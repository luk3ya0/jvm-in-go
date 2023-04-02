package references

import (
	"gopher/instruset/base"
	"gopher/rtdata"
	"gopher/rtdata/heap"
)

type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	base.InvokeMethod(frame, resolvedMethod)
}