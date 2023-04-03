package constants

import "gopher/instruset/base"
import "gopher/rtdata"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtdata.Frame) {
	// really do nothing
}
