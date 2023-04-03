package lang

import (
	"gopher/native"
	"gopher/rtdata"
	"gopher/rtdata/heap"
)

func init() {
	native.Register("java/lang/String", "intern", "()Ljava/lang/String;", intern)
}

// public native String intern();
func intern(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
