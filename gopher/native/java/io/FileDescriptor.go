package io

import "gopher/native"
import "gopher/rtdata"

const fd = "java/io/FileDescriptor"

func init() {
	native.Register(fd, "set", "(I)J", set)
}

// private static native long set(int d);
// (I)J
func set(frame *rtdata.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
