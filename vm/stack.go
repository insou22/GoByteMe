package vm

type Stack []int32

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(v int32) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int32 {
	res := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return res
}

func (s *Stack) Size() int {
	return len(*s)
}