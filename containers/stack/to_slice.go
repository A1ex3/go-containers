package stack

func (s *StackDeque[T]) ToSlice() []T {
	return s.data.ToSlice()
}

func (s *StackSlice[T]) ToSlice() []T {
	return s.data
}
