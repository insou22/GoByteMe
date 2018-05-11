package instruct

import "co.insou/gobyteme/vm"

type Stores0 struct {}
type Stores1 struct {}
type Stores2 struct {}
type Stores3 struct {}
type Stores struct {}

func (Stores0) Execute() {
	data := vm.GetStack().Pop()
	vm.GetRegister().SetData(0, data)
	vm.GetStack().Push(data)
}

func (Stores1) Execute() {
	data := vm.GetStack().Pop()
	vm.GetRegister().SetData(1, data)
	vm.GetStack().Push(data)
}

func (Stores2) Execute() {
	data := vm.GetStack().Pop()
	vm.GetRegister().SetData(2, data)
	vm.GetStack().Push(data)
}

func (Stores3) Execute() {
	data := vm.GetStack().Pop()
	vm.GetRegister().SetData(3, data)
	vm.GetStack().Push(data)
}

func (Stores) Execute() {
	register := vm.GetStack().Pop()
	data := vm.GetStack().Pop()
	vm.GetRegister().SetData(register, data)
	vm.GetStack().Push(data)
}

func (Stores0) Opcode() uint32 {
	return 37
}

func (Stores1) Opcode() uint32 {
	return 38
}

func (Stores2) Opcode() uint32 {
	return 39
}

func (Stores3) Opcode() uint32 {
	return 40
}

func (Stores) Opcode() uint32 {
	return 41
}

func (Stores0) Name() string {
	return "stores0"
}

func (Stores1) Name() string {
	return "stores1"
}

func (Stores2) Name() string {
	return "stores2"
}

func (Stores3) Name() string {
	return "stores3"
}

func (Stores) Name() string {
	return "stores"
}

func (Stores0) Length() uint32 {
	return 1
}

func (Stores1) Length() uint32 {
	return 1
}

func (Stores2) Length() uint32 {
	return 1
}

func (Stores3) Length() uint32 {
	return 1
}

func (Stores) Length() uint32 {
	return 2
}