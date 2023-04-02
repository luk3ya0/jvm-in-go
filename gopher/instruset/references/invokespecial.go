package references

import (
	"gopher/instruset/base"
	"gopher/rtdata"
)

type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtdata.Frame) {
	frame.OperandStack().PopRef()
}
