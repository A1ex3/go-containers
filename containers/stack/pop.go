package stack

import (
	"errors"
	"unsafe"
)

func (s *stack[T]) pop() (T, error) {
	l := len(s.data)
	if l == 0 {
		var t byte = 0
		return *(*T)(unsafe.Pointer(&t)), errors.New("empty stack")
	}

	// Retrieve the top element from the stack
	res := s.data[l-1]
	// Remove the top element from the stack
	s.data = s.data[:l-1]
	s.size--
	return res, nil
}

// Pop removes and returns the top element from the stack.
// If the stack is empty, it returns a zero value for T and an error.
func (s *Stack[T]) Pop() (T, error) {
	return s.stack_.pop()
}
