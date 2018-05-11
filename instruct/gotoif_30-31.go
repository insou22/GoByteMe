package instruct

import "co.insou/gobyteme/vm"

type GotoIf struct {}
type GotoSIf struct {}

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

func (GotoIf) Opcode() uint32 {
	return 30
}

func (GotoSIf) Opcode() uint32 {
	return 31
}

func (GotoIf) Name() string {
	return "gotoif"
}

func (GotoSIf) Name() string {
	return "gotosif"
}

func (GotoIf) Length() uint32 {
	return 2
}

func (GotoSIf) Length() uint32 {
	return 1
}
