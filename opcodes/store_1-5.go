package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Store0 struct {
	instruct.DefaultInstruction
}

type Store1 struct {
	instruct.DefaultInstruction
}

type Store2 struct {
	instruct.DefaultInstruction
}

type Store3 struct {
	instruct.DefaultInstruction
}

type Store struct {
	instruct.TwoByteInstruction
}

func (Store0) Execute() {
	vm.GetRegister().SetData(0, vm.GetStack().Pop())
}

func (Store1) Execute() {
	vm.GetRegister().SetData(1, vm.GetStack().Pop())
}

func (Store2) Execute() {
	vm.GetRegister().SetData(2, vm.GetStack().Pop())
}

func (Store3) Execute() {
	vm.GetRegister().SetData(3, vm.GetStack().Pop())
}

func (Store) Execute() {
	vm.GetRegister().SetData(vm.Next(), vm.GetStack().Pop())
}

func (Store0) Info() (opcode uint32, name string) {
	return 1, "store0"
}

func (Store1) Info() (opcode uint32, name string) {
	return 2, "store1"
}

func (Store2) Info() (opcode uint32, name string) {
	return 3, "store2"
}

func (Store3) Info() (opcode uint32, name string) {
	return 4, "store3"
}

func (Store) Info() (opcode uint32, name string) {
	return 5, "store"
}
