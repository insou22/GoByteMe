package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Load0 struct {
	instruct.DefaultInstruction
}

type Load1 struct {
	instruct.DefaultInstruction
}

type Load2 struct {
	instruct.DefaultInstruction
}

type Load3 struct {
	instruct.DefaultInstruction
}

type Load struct {
	instruct.TwoByteInstruction
}

func (Load0) Execute() {
	vm.GetStack().Push(vm.GetRegister().GetData(0))
}

func (Load1) Execute() {
	vm.GetStack().Push(vm.GetRegister().GetData(1))
}

func (Load2) Execute() {
	vm.GetStack().Push(vm.GetRegister().GetData(2))
}

func (Load3) Execute() {
	vm.GetStack().Push(vm.GetRegister().GetData(3))
}

func (Load) Execute() {
	vm.GetStack().Push(vm.GetRegister().GetData(vm.Next()))
}

func (Load0) Info() (opcode uint32, name string) {
	return 6, "load0"
}

func (Load1) Info() (opcode uint32, name string) {
	return 7, "load1"
}

func (Load2) Info() (opcode uint32, name string) {
	return 8, "load2"
}

func (Load3) Info() (opcode uint32, name string) {
	return 9, "load3"
}

func (Load) Info() (opcode uint32, name string) {
	return 10, "load"
}
