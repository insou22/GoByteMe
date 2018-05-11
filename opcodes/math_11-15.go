package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Sum struct {
	instruct.DefaultInstruction
}

type Sub struct {
	instruct.DefaultInstruction
}

type Mul struct {
	instruct.DefaultInstruction
}

type Div struct {
	instruct.DefaultInstruction
}

type Mod struct {
	instruct.DefaultInstruction
}

func (Sum) Execute() {
	vm.GetStack().Push(vm.GetStack().Pop() + vm.GetStack().Pop())
}

func (Sub) Execute() {
	sub := vm.GetStack().Pop()
	base := vm.GetStack().Pop()
	vm.GetStack().Push(base - sub)
}

func (Mul) Execute() {
	vm.GetStack().Push(vm.GetStack().Pop() * vm.GetStack().Pop())
}

func (Div) Execute() {
	div := vm.GetStack().Pop()
	base := vm.GetStack().Pop()
	vm.GetStack().Push(base / div)
}

func (Mod) Execute() {
	mod := vm.GetStack().Pop()
	base := vm.GetStack().Pop()
	vm.GetStack().Push(base % mod)
}

func (Sum) Info() (opcode uint32, name string) {
	return 11, "sum"
}

func (Sub) Info() (opcode uint32, name string) {
	return 12, "sub"
}

func (Mul) Info() (opcode uint32, name string) {
	return 13, "mul"
}

func (Div) Info() (opcode uint32, name string) {
	return 14, "div"
}

func (Mod) Info() (opcode uint32, name string) {
	return 15, "mod"
}
