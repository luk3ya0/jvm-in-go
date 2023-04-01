package control

import (
	"gopher/instruset/base"
	"gopher/rtdata"
)

type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtdata.Frame) {
	base.Branch(frame, self.Offset)
}
