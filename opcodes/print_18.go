package opcodes

import (
	"co.insou/gobyteme/instruct"
	"co.insou/gobyteme/vm"
	"fmt"
)

type Print struct {
	instruct.TwoByteInstruction
}

func (Print) Info() (ex func(), opcode uint32, name string) {
	return func() {
		amount := int(vm.Next())

		for i := 0; i < amount; i++ {
			fmt.Print(vm.GetStack().Pop())
		}
	},18, "print"
}
