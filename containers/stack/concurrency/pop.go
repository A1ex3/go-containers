package concurrency

func (s *stackConcurrency[T]) pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.stackSync_.Pop()
}

// Pop removes and returns the top element from the stack.
// It locks the stack to ensure thread-safe access.
// If the stack is empty, it returns an error.
func (s *StackConcurrency[T]) Pop() (T, error) {
	return s.stack_.pop()
}
