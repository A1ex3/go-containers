package linkedlist

import "errors"

func (ll *linkedList[T]) find(predicate func(T) bool) (T, error) {
	for current := ll.head; current != nil; current = current.next {
		if predicate(current.data) {
			return current.data, nil
		}
	}
	var zero T
	return zero, errors.New("no element found")
}

func (ll *LinkedList[T]) Find(predicate func(T) bool) (T, error) {
	return ll.linkedList_.find(predicate)
}
