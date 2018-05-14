package instruct

type Instruction interface {

	Info() (ex func(), opcode uint32, name string)
	Length() uint32

}

type DefaultInstruction struct {}

func (DefaultInstruction) Length() uint32 {
	return 1
}

type TwoByteInstruction struct {}

func (TwoByteInstruction) Length() uint32 {
	return 2
}