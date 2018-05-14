package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type And struct {
	instruct.DefaultInstruction
}

func (And) Info() (ex func(), opcode uint32, name string) {
	return func() {
		a1 := vm.GetStack().Pop()
		a2 := vm.GetStack().Pop()

		vm.GetStack().Push(vm.Bool2Int(a1 != 0 && a2 != 0))
	}, 42, "and"
}


type Or struct {
	instruct.DefaultInstruction
}

func (Or) Info() (ex func(), opcode uint32, name string) {
	return func() {
		a1 := vm.GetStack().Pop()
		a2 := vm.GetStack().Pop()

		vm.GetStack().Push(vm.Bool2Int(a1 != 0 || a2 != 0))
	}, 43, "or"
}


type Xor struct {
	instruct.DefaultInstruction
}

func (Xor) Info() (ex func(), opcode uint32, name string) {
	return func() {
		a1 := vm.Int2Bool(vm.GetStack().Pop())
		a2 := vm.Int2Bool(vm.GetStack().Pop())

		vm.GetStack().Push(vm.Bool2Int(a1 != a2))
	},44, "xor"
}
