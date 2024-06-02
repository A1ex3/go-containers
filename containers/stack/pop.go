package stack

import "errors"

// Pop removes and returns the top element from the stack.
// If the stack is empty, it returns a zero value for T and an error.
func (s *StackDeque[T]) Pop() (T, error) {
	return s.data.PopLast()
}

// Pop removes and returns the top element from the stack.
// If the stack is empty, it returns a zero value for T and an error.
func (s *StackSlice[T]) Pop() (T, error) {
	if s.size == 0 {
		var zero T
		return zero, errors.New("empty stack")
	}

	s.size--
	res := s.data[s.size]
	s.data[s.size] = *new(T)
	return res, nil
}
