package goregvm

type Stack struct {
	inner_stack []int32
	sp          int
}

// Push method pushes the val into the top of stack.
func (s *Stack) Push(val int32) {
	// fmt.Println("Adding ", val)
	s.inner_stack = append(s.inner_stack, val)
	s.sp++
}

// Pop method, pops the top of stack into the reg
// and decrements the stack-pointer by one.
func (s *Stack) Pop(reg Register) {

	// fmt.Println("Pop:", s.inner_stack[s.sp])
	reg.Set(s.inner_stack[s.sp])

	s.sp--

}

// GetSP method returns the address of top of stack.
func (s *Stack) GetSP() int32 {
	return int32(s.sp)
}

func NewStack() *Stack {
	return &Stack{sp: -1}
}
