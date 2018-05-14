package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type GotoIf struct {
	instruct.TwoByteInstruction
}

func (GotoIf) Info() (ex func(), opcode uint32, name string) {
	return func() {
		index := vm.Next()

		if vm.Int2Bool(vm.GetStack().Pop()) {
			vm.SetPC(uint32(index))
		}
	},30, "gotoif"
}


type GotoSIf struct {
	instruct.DefaultInstruction
}

func (GotoSIf) Info() (ex func(), opcode uint32, name string) {
	return func() {
		cmp := vm.GetStack().Pop()
		loc := vm.GetStack().Pop()

		if vm.Int2Bool(cmp) {
			vm.SetPC(uint32(loc))
		}
	},31, "gotosif"
}
