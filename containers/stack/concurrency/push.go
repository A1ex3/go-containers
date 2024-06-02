package concurrency

// Push adds a new element to the top of the stack.
// It locks the stack to ensure thread-safe access.
func (s *StackDequeConcurrency[T]) Push(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.stackDeque_.Push(value)
}

func (s *StackSliceConcurrency[T]) Push(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.stackSlice_.Push(value)
}
