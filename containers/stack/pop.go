package stack

import (
	"errors"
	"unsafe"
)

func (s *stack[T]) pop() (T, error) {
	if s.data.Size() == 0 {
		var t byte = 0
		return *(*T)(unsafe.Pointer(&t)), errors.New("empty stack")
	}

	// Retrieve the top element from the stack
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
// If the stack is empty, it returns a zero value for T and an error.
func (s *Stack[T]) Pop() (T, error) {
	return s.stack_.pop()
}
