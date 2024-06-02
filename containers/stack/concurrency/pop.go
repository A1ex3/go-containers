package concurrency

// Pop removes and returns the top element from the stack.
// It locks the stack to ensure thread-safe access.
// If the stack is empty, it returns an error.
func (s *StackDequeConcurrency[T]) Pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.stackDeque_.Pop()
}

func (s *StackSliceConcurrency[T]) Pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.stackSlice_.Pop()
}
