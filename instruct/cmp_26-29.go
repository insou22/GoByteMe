package instruct

import "co.insou/gobyteme/vm"

type CmpGt struct {}
type CmpGe struct {}
type CmpLt struct {}
type CmpLe struct {}

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

func (CmpGt) Opcode() uint32 {
	return 26
}

func (CmpGe) Opcode() uint32 {
	return 27
}

func (CmpLt) Opcode() uint32 {
	return 28
}

func (CmpLe) Opcode() uint32 {
	return 29
}

func (CmpGt) Name() string {
	return "cmpgt"
}

func (CmpGe) Name() string {
	return "cmpge"
}

func (CmpLt) Name() string {
	return "cmplt"
}

func (CmpLe) Name() string {
	return "cmple"
}

func (CmpGt) Length() uint32 {
	return 1
}

func (CmpGe) Length() uint32 {
	return 1
}

func (CmpLt) Length() uint32 {
	return 1
}

func (CmpLe) Length() uint32 {
	return 1
}

