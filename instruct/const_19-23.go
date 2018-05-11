package instruct

import "co.insou/gobyteme/vm"

type Const0 struct {}
type Const1 struct {}
type Const2 struct {}
type Const3 struct {}
type Const struct {}

func (Const0) Execute() {
	vm.GetStack().Push(0)
}

func (Const1) Execute() {
	vm.GetStack().Push(1)
}

func (Const2) Execute() {
	vm.GetStack().Push(2)
}

func (Const3) Execute() {
	vm.GetStack().Push(3)
}

func (Const) Execute() {
	vm.GetStack().Push(vm.Next())
}

func (Const0) Opcode() uint32 {
	return 19
}

func (Const1) Opcode() uint32 {
	return 20
}

func (Const2) Opcode() uint32 {
	return 21
}

func (Const3) Opcode() uint32 {
	return 22
}

func (Const) Opcode() uint32 {
	return 23
}

func (Const0) Name() string {
	return "const0"
}

func (Const1) Name() string {
	return "const1"
}

func (Const2) Name() string {
	return "const2"
}

func (Const3) Name() string {
	return "const3"
}

func (Const) Name() string {
	return "const"
}

func (Const0) Length() uint32 {
	return 1
}

func (Const1) Length() uint32 {
	return 1
}

func (Const2) Length() uint32 {
	return 1
}

func (Const3) Length() uint32 {
	return 1
}

func (Const) Length() uint32 {
	return 2
}
