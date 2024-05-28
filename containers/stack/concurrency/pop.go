package concurrency

import (
	"errors"
	"unsafe"
)

func (s *stackConcurrency[T]) pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.size == 0 {
		// Returning a zero value for T when the stack is empty
		var t byte = 0
		return *(*T)(unsafe.Pointer(&t)), errors.New("empty stack")
	}

	res := s.data[s.size-1]
	// Remove the top element from the stack
	s.data = s.data[:s.size-1]
	s.size--
	return res, nil
}

// Pop removes and returns the top element from the stack.
// It locks the stack to ensure thread-safe access.
// If the stack is empty, it returns an error.
func (s *StackConcurrency[T]) Pop() (T, error) {
	return s.stack_.pop()
}
