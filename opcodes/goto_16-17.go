package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Goto struct {
	instruct.TwoByteInstruction
}

type GotoS struct {
	instruct.DefaultInstruction
}

func (Goto) Execute() {
	vm.SetPC(uint32(vm.Next()))
}

func (GotoS) Execute() {
	vm.SetPC(uint32(vm.GetStack().Pop()))
}

func (Goto) Info() (opcode uint32, name string) {
	return 16, "goto"
}

func (GotoS) Info() (opcode uint32, name string) {
	return 17, "gotos"
}
