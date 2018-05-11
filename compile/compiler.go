package compile

import (
	"strconv"
	"co.insou/gobyteme/oplookup"
	"co.insou/gobyteme/instruct"
)

func Compile(parts []string) []int32 {
	instruct.InitialzeInstructions()

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

		inst := oplookup.InstructionFromName(val)
		ignore += int32(inst.Length() - 1)
		code[i] = int32(inst.Opcode())
	}

	return code
}