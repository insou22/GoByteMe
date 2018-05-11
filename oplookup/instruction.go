package oplookup

type Instruction interface {

	Execute()
	Opcode() uint32
	Name() string
	Length() uint32

}
