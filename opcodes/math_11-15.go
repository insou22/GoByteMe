package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Sum struct {
	instruct.DefaultInstruction
}

func (Sum) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(vm.GetStack().Pop() + vm.GetStack().Pop())
	},11, "sum"
}


type Sub struct {
	instruct.DefaultInstruction
}

func (Sub) Info() (ex func(), opcode uint32, name string) {
	return func() {
		sub := vm.GetStack().Pop()
		base := vm.GetStack().Pop()
		vm.GetStack().Push(base - sub)
	},12, "sub"
}


type Mul struct {
	instruct.DefaultInstruction
}

func (Mul) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(vm.GetStack().Pop() * vm.GetStack().Pop())
	},13, "mul"
}


type Div struct {
	instruct.DefaultInstruction
}

func (Div) Info() (ex func(), opcode uint32, name string) {
	return func() {
		div := vm.GetStack().Pop()
		base := vm.GetStack().Pop()
		vm.GetStack().Push(base / div)
	},14, "div"
}


type Mod struct {
	instruct.DefaultInstruction
}

func (Mod) Info() (ex func(), opcode uint32, name string) {
	return func() {
		mod := vm.GetStack().Pop()
		base := vm.GetStack().Pop()
		vm.GetStack().Push(base % mod)
	},15, "mod"
}
