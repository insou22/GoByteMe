package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Goto struct {
	instruct.TwoByteInstruction
}

func (Goto) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.SetPC(uint32(vm.Next()))
	},16, "goto"
}


type GotoS struct {
	instruct.DefaultInstruction
}

func (GotoS) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.SetPC(uint32(vm.GetStack().Pop()))
	},17, "gotos"
}
