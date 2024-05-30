package linkedlist

import "reflect"

func (ll *linkedList[T]) contains(value T) bool {
	current := ll.head
	for current != nil {
		if reflect.DeepEqual(current.data, value) {
			return true
		}
		current = current.next
	}
	return false
}

// Contains checks whether the specified value exists in the linked list.
// It iterates through each node and compares the node's data with the given value using reflect.DeepEqual.
// It returns true if the value is found, otherwise false.
func (ll *LinkedList[T]) Contains(value T) bool {
	return ll.linkedList_.contains(value)
}
