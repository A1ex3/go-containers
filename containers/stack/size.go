package stack

// Size returns the current number of elements in the stack.
func (s *StackDeque[T]) Size() int {
	return s.data.Size()
}

// Size returns the current number of elements in the stack.
func (s *StackSlice[T]) Size() int {
	return s.size
}
