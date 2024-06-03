package linkedlist

import "errors"

type LinkedListIterator[T any] struct {
	current *LinkedListNode[T]
}

func (linkedList *LinkedList[T]) NewIterator() *LinkedListIterator[T] {
	return &LinkedListIterator[T]{current: linkedList.linkedList_.head}
}

func (it *LinkedListIterator[T]) HasNext() bool {
	return it.current != nil
}

func (it *LinkedListIterator[T]) Next() (T, error) {
	if it.current == nil {
		var zero T
		return zero, errors.New("no more elements")
	}
	data := it.current.data
	it.current = it.current.next
	return data, nil
}
