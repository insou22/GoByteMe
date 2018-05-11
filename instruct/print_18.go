package instruct

import (
	"co.insou/gobyteme/vm"
	"fmt"
)

type Print struct {}

func (Print) Execute() {
	amount := int(vm.Next())

	for i := 0; i < amount; i++ {
		fmt.Print(vm.GetStack().Pop())
	}
}

func (Print) Opcode() uint32 {
	return 18
}

func (Print) Name() string {
	return "print"
}

func (Print) Length() uint32 {
	return 2
}
