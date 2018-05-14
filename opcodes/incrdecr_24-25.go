package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Incr struct {
	instruct.DefaultInstruction
}

func (Incr) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(vm.GetStack().Pop() + 1)
	},24, "incr"
}


type Decr struct {
	instruct.DefaultInstruction
}

func (Decr) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(vm.GetStack().Pop() - 1)
	},25, "decr"
}
