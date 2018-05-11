package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type And struct {
	instruct.DefaultInstruction
}

type Or struct {
	instruct.DefaultInstruction
}

type Xor struct {
	instruct.DefaultInstruction
}

func (And) Execute() {
	a1 := vm.GetStack().Pop()
	a2 := vm.GetStack().Pop()

	vm.GetStack().Push(vm.Bool2Int(a1 != 0 && a2 != 0))
}

func (Or) Execute() {
	a1 := vm.GetStack().Pop()
	a2 := vm.GetStack().Pop()

	vm.GetStack().Push(vm.Bool2Int(a1 != 0 || a2 != 0))
}

func (Xor) Execute() {
	a1 := vm.Int2Bool(vm.GetStack().Pop())
	a2 := vm.Int2Bool(vm.GetStack().Pop())

	vm.GetStack().Push(vm.Bool2Int(a1 != a2))
}

func (And) Info() (opcode uint32, name string) {
	return 42, "and"
}

func (Or) Info() (opcode uint32, name string) {
	return 43, "or"
}

func (Xor) Info() (opcode uint32, name string) {
	return 44, "xor"
}
