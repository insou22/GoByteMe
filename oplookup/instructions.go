package oplookup

var opcode2Inst = make(map[uint32]Instruction)
var name2Inst = make(map[string]Instruction)

func InstructionFromOpcode(opcode uint32) Instruction {
	return opcode2Inst[opcode]
}

func InstructionFromName(name string) Instruction {
	return name2Inst[name]
}

func Initialise(instruction Instruction) {
	opcode2Inst[instruction.Opcode()] = instruction
	name2Inst[instruction.Name()] = instruction
}
