package stack

import "errors"

func (s *StackDeque[T]) Top() (T, error) {
	lastElem, err := s.data.GetLast()
	if err != nil {
		return lastElem, errors.New("stack is empty")
	} else {
		return lastElem, nil
	}
}

func (s *StackSlice[T]) Top() (T, error) {
	if s.size == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	} else {
		return s.data[s.size-1], nil
	}
}
