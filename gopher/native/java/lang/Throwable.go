package lang

import (
	"fmt"
	"gopher/native"
	"gopher/rtdata"
	"gopher/rtdata/heap"
)

type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace",
		"(I)Ljava/lang/Throwable;", fillInStackTrace)
}

func (self *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		self.className, self.methodName, self.fileName, self.lineNumber)
}

// private native Throwable fillInStackTrace (int dummy);
func fillInStackTrace(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)
	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *rtdata.Thread) []*StackTraceElement {
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}

	return stes
}

func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rtdata.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}
