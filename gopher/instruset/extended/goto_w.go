package extended

import "gopher/instruset/base"
import "gopher/rtdata"

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtdata.Frame) {
	base.Branch(frame, self.offset)
}
