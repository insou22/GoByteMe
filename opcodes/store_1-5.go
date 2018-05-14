package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Store0 struct {
	instruct.DefaultInstruction
}

func (Store0) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetRegister().SetData(0, vm.GetStack().Pop())
	},1, "store0"
}


type Store1 struct {
	instruct.DefaultInstruction
}

func (Store1) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetRegister().SetData(1, vm.GetStack().Pop())
	},2, "store1"
}


type Store2 struct {
	instruct.DefaultInstruction
}

func (Store2) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetRegister().SetData(2, vm.GetStack().Pop())
	},3, "store2"
}


type Store3 struct {
	instruct.DefaultInstruction
}

func (Store3) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetRegister().SetData(3, vm.GetStack().Pop())
	},4, "store3"
}


type Store struct {
	instruct.TwoByteInstruction
}

func (Store) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetRegister().SetData(vm.Next(), vm.GetStack().Pop())
	},5, "store"
}
