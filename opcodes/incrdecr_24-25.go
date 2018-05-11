package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Incr struct {
	instruct.DefaultInstruction
}

type Decr struct {
	instruct.DefaultInstruction
}

func (Incr) Execute() {
	vm.GetStack().Push(vm.GetStack().Pop() + 1)
}

func (Decr) Execute() {
	vm.GetStack().Push(vm.GetStack().Pop() - 1)
}

func (Incr) Info() (opcode uint32, name string) {
	return 24, "incr"
}

func (Decr) Info() (opcode uint32, name string) {
	return 25, "decr"
}
