package control

import "gopher/instruset/base"
import "gopher/rtdata"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtdata.Frame) {
	base.Branch(frame, self.Offset)
}
