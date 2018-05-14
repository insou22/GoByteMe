package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Equal struct {
	instruct.DefaultInstruction
}

func (Equal) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(vm.Bool2Int(vm.GetStack().Pop() == vm.GetStack().Pop()))
	},32, "equal"
}


type Not struct {
	instruct.DefaultInstruction
}

func (Not) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(vm.Bool2Int(!vm.Int2Bool(vm.GetStack().Pop())))
	},33, "not"
}
