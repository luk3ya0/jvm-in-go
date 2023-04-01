package math

import (
	"gopher/instruset/base"
	"gopher/rtdata"
)

type ISHL struct{ base.NoOperandsInstruction }

func (self *ISHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopInt()
	v1 := stack.PopInt()

	s := uint32(v2) & 0x1f // 0b00011111 = 2^5 = 32 and s must be unsigned
	result := v1 << s
	stack.PushInt(result)
}

type ISHR struct{ base.NoOperandsInstruction }

func (self *ISHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopInt()
	v1 := stack.PopInt()

	s := uint32(v2) & 0x1f
	result := v1 >> s

	stack.PushInt(result)
}

type IUSHR struct{ base.NoOperandsInstruction }

func (self *IUSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopInt()
	v1 := stack.PopInt()

	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

type LSHL struct{ base.NoOperandsInstruction }

func (self *LSHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopInt()
	v1 := stack.PopLong()

	s := uint32(v2) & 0x3f
	result := v1 << s

	stack.PushLong(result)
}

type LSHR struct{ base.NoOperandsInstruction }

func (self *LSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopInt()
	v1 := stack.PopLong()

	s := uint32(v2) & 0x3f
	result := v1 >> s

	stack.PushLong(result)
}

type LUSHR struct{ base.NoOperandsInstruction }

func (self *LUSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopInt()
	v1 := stack.PopLong()

	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
