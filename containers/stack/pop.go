package stack

import "errors"

// Pop removes and returns the top element from the stack.
// If the stack is empty, it returns a zero value for T and an error.
func (s *StackDeque[T]) Pop() (T, error) {
	return s.Data.PopLast()
}

// Pop removes and returns the top element from the stack.
// If the stack is empty, it returns a zero value for T and an error.
func (s *StackSlice[T]) Pop() (T, error) {
	if s.Size_ == 0 {
		var zero T
		return zero, errors.New("empty stack")
	}

	s.Size_--
	res := s.Data[s.Size_]
	s.Data[s.Size_] = *new(T)
	return res, nil
}
