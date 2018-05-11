package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Exit struct {
	instruct.DefaultInstruction
}

func (Exit) Execute() {
	vm.SetResult(vm.ResultSuccessExit)
	vm.SetRunning(false)
}

func (Exit) Info() (opcode uint32, name string) {
	return 0, "exit"
}
