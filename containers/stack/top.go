package stack

import "errors"

func (s *stack[T]) top() (T, error) {
	lastElem, err := s.data.GetLast()
	if err != nil {
		return lastElem, errors.New("stack is empty")
	} else {
		return lastElem, nil
	}
}

func (s *Stack[T]) Top() (T, error) {
	return s.stack_.top()
}
