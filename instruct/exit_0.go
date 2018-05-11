package instruct

import "co.insou/gobyteme/vm"

type Exit struct {}

func (Exit) Execute() {
	vm.SetResult(vm.ResultSuccessExit)
	vm.SetRunning(false)
}

func (Exit) Opcode() uint32 {
	return 0
}

func (Exit) Name() string {
	return "exit"
}

func (Exit) Length() uint32 {
	return 1
}
