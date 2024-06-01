package linkedlist

import "errors"

func (ll *linkedList[T]) getLast() (T, error) {
	if tail := ll.tail; tail == nil {
		var zero T
		return zero, errors.New("list is empty")
	} else {
		return tail.data, nil
	}
}

func (ll *LinkedList[T]) GetLast() (T, error) {
	return ll.linkedList_.getLast()
}
