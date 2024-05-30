package linkedlist

import (
	"errors"
	"unsafe"
)

func (ll *linkedList[T]) getByIndex(index int) (T, error) {
	if index < 0 || index >= ll.size {
		var t byte = 0
		return *(*T)(unsafe.Pointer(&t)), errors.New("index out of range")
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
