package stack

import "errors"

func (s *StackDeque[T]) Top() (T, error) {
	lastElem, err := s.Data.GetLast()
	if err != nil {
		return lastElem, errors.New("stack is empty")
	} else {
		return lastElem, nil
	}
}

func (s *StackSlice[T]) Top() (T, error) {
	if s.Size_ == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	} else {
		return s.Data[s.Size_-1], nil
	}
}
