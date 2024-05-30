package concurrency

func (s *stackConcurrency[T]) size_() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data.Size()
}

// Size returns the current number of elements in the stack.
// It locks the stack to ensure thread-safe access.
func (s *StackConcurrency[T]) Size() int {
	return s.stack_.size_()
}
