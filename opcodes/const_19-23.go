package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Const0 struct {
	instruct.DefaultInstruction
}

type Const1 struct {
	instruct.DefaultInstruction
}

type Const2 struct {
	instruct.DefaultInstruction
}

type Const3 struct {
	instruct.DefaultInstruction
}

type Const struct {
	instruct.TwoByteInstruction
}

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

func (Const0) Info() (opcode uint32, name string) {
	return 19, "const0"
}

func (Const1) Info() (opcode uint32, name string) {
	return 20, "const1"
}

func (Const2) Info() (opcode uint32, name string) {
	return 21, "const2"
}

func (Const3) Info() (opcode uint32, name string) {
	return 22, "const3"
}

func (Const) Info() (opcode uint32, name string) {
	return 23, "const"
}
