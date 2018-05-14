package instruct

var opcode2Inst = make(map[uint32]Instruction)
var name2Inst = make(map[string]Instruction)

func InstructionFromOpcode(opcode uint32) Instruction {
	return opcode2Inst[opcode]
}

func InstructionFromName(name string) Instruction {
	return name2Inst[name]
}

func Initialise(instruction Instruction) {
	_, opcode, name := instruction.Info()
	opcode2Inst[opcode] = instruction
	name2Inst[name] = instruction
}
