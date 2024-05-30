package concurrency

import (
	"errors"
	"unsafe"
)

func (s *stackConcurrency[T]) pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.data.Size() == 0 {
		// Returning a zero value for T when the stack is empty
		var t byte = 0
		return *(*T)(unsafe.Pointer(&t)), errors.New("empty stack")
	}

	res, err := s.data.GetByIndex(s.data.Size() - 1)
	if err != nil {
		var t byte = 0
		return *(*T)(unsafe.Pointer(&t)), err
	}
	// Remove the top element from the stack
	if !s.data.Remove(s.data.Size() - 1) {
		var t byte = 0
		return *(*T)(unsafe.Pointer(&t)), errors.New("failed to remove an element from the stack")
	}
	return res, nil
}

// Pop removes and returns the top element from the stack.
// It locks the stack to ensure thread-safe access.
// If the stack is empty, it returns an error.
func (s *StackConcurrency[T]) Pop() (T, error) {
	return s.stack_.pop()
}
