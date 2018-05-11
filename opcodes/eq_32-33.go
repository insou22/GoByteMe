package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Equal struct {
	instruct.DefaultInstruction
}

type Not struct {
	instruct.DefaultInstruction
}

func (Equal) Execute() {
	vm.GetStack().Push(vm.Bool2Int(vm.GetStack().Pop() == vm.GetStack().Pop()))
}

func (Not) Execute() {
	vm.GetStack().Push(vm.Bool2Int(!vm.Int2Bool(vm.GetStack().Pop())))
}

func (Equal) Info() (opcode uint32, name string) {
	return 32, "equal"
}

func (Not) Info() (opcode uint32, name string) {
	return 33, "not"
}
