package compile

import (
	"strconv"
	"co.insou/gobyteme/instruct"
	"co.insou/gobyteme/opcodes"
)

func Compile(parts []string) []int32 {
	opcodes.InitialzeInstructions()

	ignore := int32(0)

	code := make([]int32, len(parts))

	for i := range parts {
		val := parts[i]

		if ignore > 0 {
			ignore--
			op, _ := strconv.Atoi(val)
			code[i] = int32(op)
			continue
		}

		inst := instruct.InstructionFromName(val)
		ignore += int32(inst.Length() - 1)
		_, opcode, _ := inst.Info()
		code[i] = int32(opcode)
	}

	return code
}