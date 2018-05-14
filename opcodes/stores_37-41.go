package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Stores0 struct {
	instruct.DefaultInstruction
}

func (Stores0) Info() (ex func(), opcode uint32, name string) {
	return func() {
		data := vm.GetStack().Pop()
		vm.GetRegister().SetData(0, data)
		vm.GetStack().Push(data)
	},37, "stores0"
}


type Stores1 struct {
	instruct.DefaultInstruction
}

func (Stores1) Info() (ex func(), opcode uint32, name string) {
	return func() {
		data := vm.GetStack().Pop()
		vm.GetRegister().SetData(1, data)
		vm.GetStack().Push(data)
	},38, "stores1"
}


type Stores2 struct {
	instruct.DefaultInstruction
}

func (Stores2) Info() (ex func(), opcode uint32, name string) {
	return func() {
		data := vm.GetStack().Pop()
		vm.GetRegister().SetData(2, data)
		vm.GetStack().Push(data)
	},39, "stores2"
}


type Stores3 struct {
	instruct.DefaultInstruction
}

func (Stores3) Info() (ex func(), opcode uint32, name string) {
	return func() {
		data := vm.GetStack().Pop()
		vm.GetRegister().SetData(3, data)
		vm.GetStack().Push(data)
	},40, "stores3"
}


type Stores struct {
	instruct.TwoByteInstruction
}

func (Stores) Info() (ex func(), opcode uint32, name string) {
	return func() {
		register := vm.GetStack().Pop()
		data := vm.GetStack().Pop()
		vm.GetRegister().SetData(register, data)
		vm.GetStack().Push(data)
	},41, "stores"
}
