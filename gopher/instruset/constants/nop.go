package constants

import (
	"gopher/instruset/base"
	"gopher/rtdata"
)

// NOP Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtdata.Frame) { /* do nothing */ }
