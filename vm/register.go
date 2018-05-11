package vm

type Register struct {}

var reg = make(map[int32]int32)

func NewRegister() Register {
	return *new(Register)
}

func (r Register) GetData(index int32) int32 {
	return reg[index]
}

func (r Register) SetData(index, value int32) {
	reg[index] = value
}