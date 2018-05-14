package opcodes

import (
	"co.insou/gobyteme/vm"
	"fmt"
	"co.insou/gobyteme/instruct"
)

type PrintA struct {
	instruct.TwoByteInstruction
}

func (PrintA) Info() (ex func(), opcode uint32, name string) {
	return func() {
		amount := int(vm.Next())

		for i := 0; i < amount; i++ {
			fmt.Print(rune(vm.GetStack().Pop()))
		}
	},34, "printa"
}


type PrintAS struct {
	instruct.DefaultInstruction
}

func (PrintAS) Info() (ex func(), opcode uint32, name string) {
	return func() {
		amount := int(vm.GetStack().Pop())

		for i := 0; i < amount; i++ {
			fmt.Print(rune(vm.GetStack().Pop()))
		}
	},35, "printas"
}


type Print1 struct {
	instruct.DefaultInstruction
}

func (Print1) Info() (ex func(), opcode uint32, name string) {
	return func() {
		fmt.Println(vm.GetStack().Pop())
	},36, "print1"
}
