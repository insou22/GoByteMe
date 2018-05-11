package instruct

import "co.insou/gobyteme/vm"

type Store0 struct {}
type Store1 struct {}
type Store2 struct {}
type Store3 struct {}
type Store struct {}

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

func (Store0) Opcode() uint32 {
	return 1
}

func (Store1) Opcode() uint32 {
	return 2
}

func (Store2) Opcode() uint32 {
	return 3
}

func (Store3) Opcode() uint32 {
	return 4
}

func (Store) Opcode() uint32 {
	return 5
}

func (Store0) Name() string {
	return "store0"
}

func (Store1) Name() string {
	return "store1"
}

func (Store2) Name() string {
	return "store2"
}

func (Store3) Name() string {
	return "store3"
}

func (Store) Name() string {
	return "store"
}

func (Store0) Length() uint32 {
	return 1
}

func (Store1) Length() uint32 {
	return 1
}

func (Store2) Length() uint32 {
	return 1
}

func (Store3) Length() uint32 {
	return 1
}

func (Store) Length() uint32 {
	return 2
}
