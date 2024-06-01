package linkedlist

import "errors"

func (ll *linkedList[T]) getFirst() (T, error) {
	if head := ll.head; head == nil {
		var zero T
		return zero, errors.New("list is empty")
	} else {
		return head.data, nil
	}
}

func (ll *LinkedList[T]) GetFirst() (T, error) {
	return ll.linkedList_.getFirst()
}
