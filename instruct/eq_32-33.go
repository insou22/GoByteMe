package instruct

import "co.insou/gobyteme/vm"

type Equal struct {}
type Not struct {}

func (Equal) Execute() {
	vm.GetStack().Push(vm.Bool2Int(vm.GetStack().Pop() == vm.GetStack().Pop()))
}

func (Not) Execute() {
	vm.GetStack().Push(vm.Bool2Int(!vm.Int2Bool(vm.GetStack().Pop())))
}

func (Equal) Opcode() uint32 {
	return 32
}

func (Not) Opcode() uint32 {
	return 33
}

func (Equal) Name() string {
	return "equal"
}

func (Not) Name() string {
	return "not"
}

func (Equal) Length() uint32 {
	return 1
}

func (Not) Length() uint32 {
	return 1
}
