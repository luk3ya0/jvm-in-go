package stack

import "gopher/instruset/base"
import "gopher/rtdata"

// Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
