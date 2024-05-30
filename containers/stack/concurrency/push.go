package concurrency

func (s *stackConcurrency[T]) push(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.stackSync_.Push(value)
}

// Push adds a new element to the top of the stack.
// It locks the stack to ensure thread-safe access.
func (s *StackConcurrency[T]) Push(value T) {
	s.stack_.push(value)
}
