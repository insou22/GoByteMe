package instruct

import "co.insou/gobyteme/vm"

type Sum struct {}
type Sub struct {}
type Mul struct {}
type Div struct {}
type Mod struct {}

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

func (Sum) Opcode() uint32 {
	return 11
}

func (Sub) Opcode() uint32 {
	return 12
}

func (Mul) Opcode() uint32 {
	return 13
}

func (Div) Opcode() uint32 {
	return 14
}

func (Mod) Opcode() uint32 {
	return 15
}

func (Sum) Name() string {
	return "sum"
}

func (Sub) Name() string {
	return "sub"
}

func (Mul) Name() string {
	return "mul"
}

func (Div) Name() string {
	return "div"
}

func (Mod) Name() string {
	return "mod"
}

func (Sum) Length() uint32 {
	return 1
}

func (Sub) Length() uint32 {
	return 1
}

func (Mul) Length() uint32 {
	return 1
}

func (Div) Length() uint32 {
	return 1
}

func (Mod) Length() uint32 {
	return 2
}
