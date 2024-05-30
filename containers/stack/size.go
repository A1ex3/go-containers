package stack

// Size returns the current number of elements in the stack.
func (s *stack[T]) size_() int {
	return s.data.Size()
}

// Size returns the current number of elements in the stack.
func (s *Stack[T]) Size() int {
	return s.stack_.size_()
}
