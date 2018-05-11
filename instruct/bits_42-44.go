package instruct

import (
	"co.insou/gobyteme/vm"
)

type And struct {}
type Or struct {}
type Xor struct {}

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

func (And) Opcode() uint32 {
	return 42
}

func (Or) Opcode() uint32 {
	return 43
}

func (Xor) Opcode() uint32 {
	return 44
}

func (And) Name() string {
	return "and"
}

func (Or) Name() string {
	return "or"
}

func (Xor) Name() string {
	return "xor"
}

func (And) Length() uint32 {
	return 1
}

func (Or) Length() uint32 {
	return 1
}

func (Xor) Length() uint32 {
	return 1
}

