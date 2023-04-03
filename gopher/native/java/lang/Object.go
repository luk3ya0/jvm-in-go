package lang

import "gopher/native"
import "gopher/rtdata"

const jlObject = "java/lang/Object"

func init() {
	native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
}

// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(frame *rtdata.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
