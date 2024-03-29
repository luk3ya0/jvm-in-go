package reserved

import "gopher/instruset/base"
import "gopher/rtdata"
import "gopher/native"
import _ "gopher/native/java/io"
import _ "gopher/native/java/lang"
import _ "gopher/native/java/security"
import _ "gopher/native/java/util/concurrent/atomic"
import _ "gopher/native/sun/io"
import _ "gopher/native/sun/misc"
import _ "gopher/native/sun/reflect"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtdata.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
