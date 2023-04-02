package references

import (
	"gopher/instruset/base"
	"gopher/rtdata"
	"gopher/rtdata/heap"
)

type ANEW_ARRAY struct{ base.Index16Instruction }

func (self *ANEW_ARRAY) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	eleClass := classRef.ResolvedClass()
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := eleClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))

	stack.PushRef(arr)
}
