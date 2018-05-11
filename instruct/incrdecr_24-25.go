package instruct

import "co.insou/gobyteme/vm"

type Incr struct {}
type Decr struct {}

func (Incr) Execute() {
	vm.GetStack().Push(vm.GetStack().Pop() + 1)
}

func (Decr) Execute() {
	vm.GetStack().Push(vm.GetStack().Pop() - 1)
}

func (Incr) Opcode() uint32 {
	return 24
}

func (Decr) Opcode() uint32 {
	return 25
}

func (Incr) Name() string {
	return "incr"
}

func (Decr) Name() string {
	return "decr"
}

func (Incr) Length() uint32 {
	return 1
}

func (Decr) Length() uint32 {
	return 1
}
