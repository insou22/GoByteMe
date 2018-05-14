package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type Const0 struct {
	instruct.DefaultInstruction
}

func (Const0) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(0)
	},19, "const0"
}


type Const1 struct {
	instruct.DefaultInstruction
}

func (Const1) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(1)
	},20, "const1"
}


type Const2 struct {
	instruct.DefaultInstruction
}

func (Const2) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(2)
	},21, "const2"
}


type Const3 struct {
	instruct.DefaultInstruction
}

func (Const3) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(3)
	},22, "const3"
}


type Const struct {
	instruct.TwoByteInstruction
}

func (Const) Info() (ex func(), opcode uint32, name string) {
	return func() {
		vm.GetStack().Push(vm.Next())
	},23, "const"
}
