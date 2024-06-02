package stack

func (s *stack[T]) toSlice() []T {
	return s.data.ToSlice()
}

func (s *Stack[T]) ToSlice() []T {
	return s.stack_.toSlice()
}
