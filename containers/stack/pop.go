package stack

func (s *stack[T]) pop() (T, error) {
	return s.data.PopLast()
}

// Pop removes and returns the top element from the stack.
// If the stack is empty, it returns a zero value for T and an error.
func (s *Stack[T]) Pop() (T, error) {
	return s.stack_.pop()
}
