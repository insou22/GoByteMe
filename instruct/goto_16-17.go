package instruct

import "co.insou/gobyteme/vm"

type Goto struct {}
type GotoS struct {}

func (Goto) Execute() {
	vm.SetPC(uint32(vm.Next()))
}

func (GotoS) Execute() {
	vm.SetPC(uint32(vm.GetStack().Pop()))
}

func (Goto) Opcode() uint32 {
	return 16
}

func (GotoS) Opcode() uint32 {
	return 17
}

func (Goto) Name() string {
	return "goto"
}

func (GotoS) Name() string {
	return "gotos"
}

func (Goto) Length() uint32 {
	return 2
}

func (GotoS) Length() uint32 {
	return 1
}
