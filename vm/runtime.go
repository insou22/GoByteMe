package vm

import (
	"co.insou/gobyteme/oplookup"
)

var register = NewRegister()
var stack = NewStack()
var code []int32 = nil

var pc uint32 = 0
var running = false
var result uint8

const(
	ResultSuccessExit = 0
	ResultSuccessEof  = 1
	ResultFailed	  = 2
)

func LoadCode(c []int32) {
	code = c
}

func Begin() {
	if code == nil {
		panic("Code was not loaded!")
	}

	running = true

	codeLen := len(code)

	for running && int(pc) < codeLen {
		inst := oplookup.InstructionFromOpcode(uint32(Next()))
		attemptExecute(inst)
	}

	result = ResultSuccessEof
}

func attemptExecute(instruction oplookup.Instruction) {
	defer func() {
		if r := recover(); r != nil {
			result = ResultFailed
			running = false
			panic(r)
		}
	}()

	instruction.Execute()
}

func Next() int32 {
	defer func() { pc++ }()
	return code[pc]
}

func GetRegister() Register {
	return register
}

func GetStack() *Stack {
	return stack
}

func SetPC(val uint32) {
	pc = val
}

func SetRunning(run bool) {
	running = run
}

func SetResult(res uint8) {
	result = res
}

func GetResult() uint8 {
	return result
}

func Bool2Int(b bool) (r int32) {
	if b {
		r = 1
	}
	return
}

func Int2Bool(i int32) bool {
	return i != 0
}