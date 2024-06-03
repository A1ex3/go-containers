package stack

// Push adds a new element to the top of the stack.
func (s *StackDeque[T]) Push(value T) {
	s.Data.PushLast(value)
}

// Push adds a new element to the top of the stack.
func (s *StackSlice[T]) Push(value T) {
	s.Data = append(s.Data, value)
	s.Size_++
}
