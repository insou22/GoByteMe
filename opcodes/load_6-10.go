package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Load0 struct {
	instruct.DefaultInstruction
}

func (Load0) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(vm.GetRegister().GetData(0))
	},6, "load0"
}


type Load1 struct {
	instruct.DefaultInstruction
}

func (Load1) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(vm.GetRegister().GetData(1))
	},7, "load1"
}


type Load2 struct {
	instruct.DefaultInstruction
}

func (Load2) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(vm.GetRegister().GetData(2))
	},8, "load2"
}


type Load3 struct {
	instruct.DefaultInstruction
}

func (Load3) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(vm.GetRegister().GetData(3))
	},9, "load3"
}


type Load struct {
	instruct.TwoByteInstruction
}

func (Load) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(vm.GetRegister().GetData(vm.Next()))
	},10, "load"
}
