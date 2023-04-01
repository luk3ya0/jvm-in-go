package constants

import (
	"gopher/instruset/base"
	"gopher/rtdata"
)

type BIPUSH struct{ val int8 }  // Push byte
type SIPUSH struct{ val int16 } // Push short

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtdata.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtdata.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
