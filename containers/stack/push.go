package stack

// Push adds a new element to the top of the stack.
func (s *StackDeque[T]) Push(value T) {
	s.data.PushLast(value)
}

// Push adds a new element to the top of the stack.
func (s *StackSlice[T]) Push(value T) {
	s.data = append(s.data, value)
	s.size++
}
