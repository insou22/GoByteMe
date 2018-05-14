package opcodes

import (
	"co.insou/gobyteme/instruct"
	"co.insou/gobyteme/vm"
)

type Exit struct {
	instruct.DefaultInstruction
}

func (Exit) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.SetResult(vm.ResultSuccessExit)
		vm.SetRunning(false)
	},0, "exit"
}
