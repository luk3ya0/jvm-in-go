package base

import (
	"gopher/rtdata"
	"gopher/rtdata/heap"
)

func InvokeMethod(invokerFrame *rtdata.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	argSlotSlot := int(method.ArgSlotCount())
	if argSlotSlot > 0 {
		for i := argSlotSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}
