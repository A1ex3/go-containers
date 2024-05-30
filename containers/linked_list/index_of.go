package linkedlist

import (
	"errors"
	"reflect"
)

func (ll *linkedList[T]) indexOf(value T) (int, error) {
	current := ll.head
	index := 0
	for current != nil {
		if reflect.DeepEqual(current.data, value) {
			return index, nil
		}
		current = current.next
		index++
	}
	return -1, errors.New("value not found")
}

// IndexOf returns the index of the first occurrence of the specified value in the linked list.
// If the value is not found, it returns -1 along with an error.
func (ll *LinkedList[T]) IndexOf(value T) (int, error) {
	return ll.linkedList_.indexOf(value)
}
