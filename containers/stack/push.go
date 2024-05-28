package stack

func (s *stack[T]) push(value T) {
	s.data = append(s.data, value)
	s.size++
}

// Push adds a new element to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.stack_.push(value)
}
