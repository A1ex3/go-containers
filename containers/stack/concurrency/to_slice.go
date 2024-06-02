package concurrency

func (s *StackDequeConcurrency[T]) ToSlice() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stackDeque_.ToSlice()
}

func (s *StackSliceConcurrency[T]) ToSlice() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stackSlice_.ToSlice()
}
