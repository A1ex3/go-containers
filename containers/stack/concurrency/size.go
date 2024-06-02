package concurrency

// Size returns the current number of elements in the stack.
// It locks the stack to ensure thread-safe access.
func (s *StackDequeConcurrency[T]) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.stackDeque_.Size()
}

func (s *StackSliceConcurrency[T]) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.stackSlice_.Size()
}
