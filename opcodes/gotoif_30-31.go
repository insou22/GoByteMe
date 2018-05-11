package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type GotoIf struct {
	instruct.TwoByteInstruction
}

type GotoSIf struct {
	instruct.DefaultInstruction
}

func (GotoIf) Execute() {
	index := vm.Next()

	if vm.Int2Bool(vm.GetStack().Pop()) {
		vm.SetPC(uint32(index))
	}
}

func (GotoSIf) Execute() {
	cmp := vm.GetStack().Pop()
	loc := vm.GetStack().Pop()

	if vm.Int2Bool(cmp) {
		vm.SetPC(uint32(loc))
	}
}

func (GotoIf) Info() (opcode uint32, name string) {
	return 30, "gotoif"
}

func (GotoSIf) Info() (opcode uint32, name string) {
	return 31, "gotosif"
}
