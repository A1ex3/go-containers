package concurrency

func (s *stackConcurrency[T]) top() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stackSync_.Top()
}

func (s *StackConcurrency[T]) Top() (T, error) {
	return s.stack_.top()
}
