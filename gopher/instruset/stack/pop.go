package stack

import "gopher/instruset/base"
import "gopher/rtdata"

// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

func (self *POP2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
