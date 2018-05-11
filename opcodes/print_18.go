package opcodes

import (
	"co.insou/gobyteme/vm"
	"fmt"
	"co.insou/gobyteme/instruct"
)

type Print struct {
	instruct.TwoByteInstruction
}

func (Print) Execute() {
	amount := int(vm.Next())

	for i := 0; i < amount; i++ {
		fmt.Print(vm.GetStack().Pop())
	}
}

func (Print) Info() (opcode uint32, name string) {
	return 18, "print"
}
