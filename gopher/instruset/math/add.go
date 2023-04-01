package math

import (
	"gopher/instruset/base"
	"gopher/rtdata"
)

type DADD struct{ base.NoOperandsInstruction }

func (self *DADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopDouble()
	v2 := stack.PopDouble()

	result := v1 + v2
	stack.PushDouble(result)
}

type FADD struct{ base.NoOperandsInstruction }

func (self *FADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopFloat()
	v2 := stack.PopFloat()

	result := v1 + v2
	stack.PushFloat(result)
}

type IADD struct{ base.NoOperandsInstruction }

func (self *IADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopInt()
	v2 := stack.PopInt()

	result := v1 + v2
	stack.PushInt(result)
}

type LADD struct{ base.NoOperandsInstruction }

func (self *LADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopLong()
	v2 := stack.PopLong()

	result := v1 + v2
	stack.PushLong(result)
}