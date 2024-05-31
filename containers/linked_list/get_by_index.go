package linkedlist

import (
	"errors"
)

func (ll *linkedList[T]) getByIndex(index int) (T, error) {
	if index < 0 || index >= ll.size {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.data, nil
}

// GetByIndex returns the data of the node at the specified index in the linked list.
// If the index is out of bounds, it returns a zero value of type T along with an error.
func (ll *LinkedList[T]) GetByIndex(index int) (T, error) {
	return ll.linkedList_.getByIndex(index)
}
