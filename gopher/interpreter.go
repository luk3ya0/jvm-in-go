package main

import (
	"fmt"
	"gopher/instruset"
	"gopher/instruset/base"
	"gopher/rtdata"
	"gopher/rtdata/heap"
)

func interpret(method *heap.Method, logInst bool) {
	thread := rtdata.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtdata.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtdata.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()

		fmt.Printf(">> pc: %4d %v.%v %v\n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func loop(thread *rtdata.Thread, logInst bool) {
	reader := &base.BytecodeReader{}

	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instruset.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		if logInst {
			logInstruction(frame, inst)
		}

		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtdata.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()

	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}