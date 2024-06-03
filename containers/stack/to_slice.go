package stack

func (s *StackDeque[T]) ToSlice() []T {
	return s.Data.ToSlice()
}

func (s *StackSlice[T]) ToSlice() []T {
	return s.Data
}
