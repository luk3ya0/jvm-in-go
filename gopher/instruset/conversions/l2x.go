package conversions

import (
	"gopher/instruset/base"
	"gopher/rtdata"
)

type L2I struct{ base.NoOperandsInstruction }

func (self *L2I) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	l := stack.PopLong()
	i := int32(l)

	stack.PushInt(i)
}

type L2F struct{ base.NoOperandsInstruction }

func (self *L2F) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	l := stack.PopLong()
	f := float32(l)

	stack.PushFloat(f)
}

type L2D struct{ base.NoOperandsInstruction }

func (self *L2D) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	l := stack.PopLong()
	d := float64(l)

	stack.PushDouble(d)
}

