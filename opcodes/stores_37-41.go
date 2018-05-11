package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Stores0 struct {
	instruct.DefaultInstruction
}

type Stores1 struct {
	instruct.DefaultInstruction
}

type Stores2 struct {
	instruct.DefaultInstruction
}

type Stores3 struct {
	instruct.DefaultInstruction
}

type Stores struct {
	instruct.TwoByteInstruction
}

func (Stores0) Execute() {
	data := vm.GetStack().Pop()
	vm.GetRegister().SetData(0, data)
	vm.GetStack().Push(data)
}

func (Stores1) Execute() {
	data := vm.GetStack().Pop()
	vm.GetRegister().SetData(1, data)
	vm.GetStack().Push(data)
}

func (Stores2) Execute() {
	data := vm.GetStack().Pop()
	vm.GetRegister().SetData(2, data)
	vm.GetStack().Push(data)
}

func (Stores3) Execute() {
	data := vm.GetStack().Pop()
	vm.GetRegister().SetData(3, data)
	vm.GetStack().Push(data)
}

func (Stores) Execute() {
	register := vm.GetStack().Pop()
	data := vm.GetStack().Pop()
	vm.GetRegister().SetData(register, data)
	vm.GetStack().Push(data)
}

func (Stores0) Info() (opcode uint32, name string) {
	return 37, "stores0"
}

func (Stores1) Info() (opcode uint32, name string) {
	return 38, "stores1"
}

func (Stores2) Info() (opcode uint32, name string) {
	return 39, "stores2"
}

func (Stores3) Info() (opcode uint32, name string) {
	return 40, "stores3"
}

func (Stores) Info() (opcode uint32, name string) {
	return 41, "stores"
}
