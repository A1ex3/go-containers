package stack

import (
	"errors"
)

func (s *stack[T]) pop() (T, error) {
	var zero T

	if s.data.Size() == 0 {
		return zero, errors.New("empty stack")
	}

	// Retrieve the top element from the stack
	res, err := s.data.GetByIndex(s.data.Size() - 1)
	if err != nil {
		return zero, err
	}
	// Remove the top element from the stack
	if !s.data.RemoveLast() {
		return zero, errors.New("failed to remove an element from the stack")
	}
	return res, nil
}

// Pop removes and returns the top element from the stack.
// If the stack is empty, it returns a zero value for T and an error.
func (s *Stack[T]) Pop() (T, error) {
	return s.stack_.pop()
}
