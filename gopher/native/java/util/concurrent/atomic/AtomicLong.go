package atomic

import "gopher/native"
import "gopher/rtdata"

func init() {
	native.Register("java/util/concurrent/atomic/AtomicLong", "VMSupportsCS8", "()Z", vmSupportsCS8)
}

func vmSupportsCS8(frame *rtdata.Frame) {
	frame.OperandStack().PushBoolean(false)
}
