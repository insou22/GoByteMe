package opcodes

import (
	"co.insou/gobyteme/vm"
	"fmt"
	"co.insou/gobyteme/instruct"
)

type PrintA struct {
	instruct.TwoByteInstruction
}

type PrintAS struct {
	instruct.DefaultInstruction
}

type Print1 struct {
	instruct.DefaultInstruction
}

func (PrintA) Execute() {
	amount := int(vm.Next())

	for i := 0; i < amount; i++ {
		fmt.Print(rune(vm.GetStack().Pop()))
	}
}

func (PrintAS) Execute() {
	amount := int(vm.GetStack().Pop())

	for i := 0; i < amount; i++ {
		fmt.Print(rune(vm.GetStack().Pop()))
	}
}

func (Print1) Execute() {
	fmt.Println(vm.GetStack().Pop())
}

func (PrintA) Info() (opcode uint32, name string) {
	return 34, "printa"
}

func (PrintAS) Info() (opcode uint32, name string) {
	return 35, "printas"
}

func (Print1) Info() (opcode uint32, name string) {
	return 36, "print1"
}
