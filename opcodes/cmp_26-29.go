package opcodes

import (
	"co.insou/gobyteme/vm"
	"co.insou/gobyteme/instruct"
)

type CmpGt struct {
	instruct.DefaultInstruction
}

type CmpGe struct {
	instruct.DefaultInstruction
}

type CmpLt struct {
	instruct.DefaultInstruction
}

type CmpLe struct {
	instruct.DefaultInstruction
}

func (CmpGt) Execute() {
	operand := vm.GetStack().Pop()
	base := vm.GetStack().Pop()
	vm.GetStack().Push(vm.Bool2Int(base > operand))
}

func (CmpGe) Execute() {
	operand := vm.GetStack().Pop()
	base := vm.GetStack().Pop()
	vm.GetStack().Push(vm.Bool2Int(base >= operand))
}

func (CmpLt) Execute() {
	operand := vm.GetStack().Pop()
	base := vm.GetStack().Pop()
	vm.GetStack().Push(vm.Bool2Int(base < operand))
}

func (CmpLe) Execute() {
	operand := vm.GetStack().Pop()
	base := vm.GetStack().Pop()
	vm.GetStack().Push(vm.Bool2Int(base <= operand))
}

func (CmpGt) Info() (opcode uint32, name string) {
	return 26, "cmpgt"
}

func (CmpGe) Info() (opcode uint32, name string) {
	return 27, "cmpge"
}

func (CmpLt) Info() (opcode uint32, name string) {
	return 28, "cmplt"
}

func (CmpLe) Info() (opcode uint32, name string) {
	return 29, "cmple"
}
