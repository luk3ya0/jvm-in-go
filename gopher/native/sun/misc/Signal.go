package misc

import "gopher/native"
import "gopher/rtdata"

func init() {
	_signal(findSignal, "findSignal", "(Ljava/lang/String;)I")
	_signal(handle0, "handle0", "(IJ)J")
}

func _signal(method func(frame *rtdata.Frame), name, desc string) {
	native.Register("sun/misc/Signal", name, desc, method)
}

// private static native int findSignal(String string);
// (Ljava/lang/String;)I
func findSignal(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	vars.GetRef(0) // name

	stack := frame.OperandStack()
	stack.PushInt(0) // todo
}

// private static native long handle0(int i, long l);
// (IJ)J
func handle0(frame *rtdata.Frame) {
	// todo
	vars := frame.LocalVars()
	vars.GetInt(0)
	vars.GetLong(1)

	stack := frame.OperandStack()
	stack.PushLong(0)
}

// private static native void raise0(int i);