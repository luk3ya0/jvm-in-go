package lang

import "runtime"
import "gopher/native"
import "gopher/rtdata"

const jlRuntime = "java/lang/Runtime"

func init() {
	native.Register(jlRuntime, "availableProcessors", "()I", availableProcessors)
}

// public native int availableProcessors();
// ()I
func availableProcessors(frame *rtdata.Frame) {
	numCPU := runtime.NumCPU()

	stack := frame.OperandStack()
	stack.PushInt(int32(numCPU))
}
