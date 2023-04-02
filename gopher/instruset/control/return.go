package control

import (
	"gopher/instruset/base"
	"gopher/rtdata"
)

type RETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtdata.Frame) {
	frame.Thread().PopFrame()
}

type ARETURN struct{ base.NoOperandsInstruction }

func (self *ARETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()

	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()

	retVal := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retVal)
}

type DRETURN struct{ base.NoOperandsInstruction }

func (self *DRETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()

	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()

	retVal := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}

type FRETURN struct{ base.NoOperandsInstruction }

func (self *FRETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()

	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()

	retVal := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}

type IRETURN struct{ base.NoOperandsInstruction }

func (self *IRETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()

	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()

	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}

type LRETURN struct{ base.NoOperandsInstruction }

func (self *LRETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()

	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()

	retVal := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}
