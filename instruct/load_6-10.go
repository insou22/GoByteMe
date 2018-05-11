package instruct

import "co.insou/gobyteme/vm"

type Load0 struct {}
type Load1 struct {}
type Load2 struct {}
type Load3 struct {}
type Load struct {}

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

func (Load0) Opcode() uint32 {
	return 6
}

func (Load1) Opcode() uint32 {
	return 7
}

func (Load2) Opcode() uint32 {
	return 8
}

func (Load3) Opcode() uint32 {
	return 9
}

func (Load) Opcode() uint32 {
	return 10
}

func (Load0) Name() string {
	return "load0"
}

func (Load1) Name() string {
	return "load1"
}

func (Load2) Name() string {
	return "load2"
}

func (Load3) Name() string {
	return "load3"
}

func (Load) Name() string {
	return "load"
}

func (Load0) Length() uint32 {
	return 1
}

func (Load1) Length() uint32 {
	return 1
}

func (Load2) Length() uint32 {
	return 1
}

func (Load3) Length() uint32 {
	return 1
}

func (Load) Length() uint32 {
	return 2
}
