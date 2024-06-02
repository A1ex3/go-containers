package concurrency

func (s *stackConcurrency[T]) toSlice() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stackSync_.ToSlice()
}

func (s *StackConcurrency[T]) ToSlice() []T {
	return s.stack_.toSlice()
}
