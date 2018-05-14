package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type CmpGt struct {
	instruct.DefaultInstruction
}

func (CmpGt) Info() (ex func(), opcode uint32, name string) {
	return func() {
		operand := vm.GetStack().Pop()
		base := vm.GetStack().Pop()
		vm.GetStack().Push(vm.Bool2Int(base > operand))
	},26, "cmpgt"
}


type CmpGe struct {
	instruct.DefaultInstruction
}

func (CmpGe) Info() (ex func(), opcode uint32, name string) {
	return func() {
		operand := vm.GetStack().Pop()
		base := vm.GetStack().Pop()
		vm.GetStack().Push(vm.Bool2Int(base >= operand))
	},27, "cmpge"
}


type CmpLt struct {
	instruct.DefaultInstruction
}

func (CmpLt) Info() (ex func(), opcode uint32, name string) {
	return func() {
		operand := vm.GetStack().Pop()
		base := vm.GetStack().Pop()
		vm.GetStack().Push(vm.Bool2Int(base < operand))
	},28, "cmplt"
}


type CmpLe struct {
	instruct.DefaultInstruction
}

func (CmpLe) Info() (ex func(), opcode uint32, name string) {
	return func() {
		operand := vm.GetStack().Pop()
		base := vm.GetStack().Pop()
		vm.GetStack().Push(vm.Bool2Int(base <= operand))
	},29, "cmple"
}
