package concurrency

func (s *StackDequeConcurrency[T]) Top() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stackDeque_.Top()
}

func (s *StackSliceConcurrency[T]) Top() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stackSlice_.Top()
}
