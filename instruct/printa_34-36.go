package instruct

import (
	"co.insou/gobyteme/vm"
	"fmt"
)

type PrintA struct {}
type PrintAS struct {}
type Print1 struct {}

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

func (PrintA) Opcode() uint32 {
	return 34
}

func (PrintAS) Opcode() uint32 {
	return 35
}

func (Print1) Opcode() uint32 {
	return 36
}

func (PrintA) Name() string {
	return "printa"
}

func (PrintAS) Name() string {
	return "printas"
}

func (Print1) Name() string {
	return "print1"
}

func (PrintA) Length() uint32 {
	return 2
}

func (PrintAS) Length() uint32 {
	return 1
}

func (Print1) Length() uint32 {
	return 1
}
