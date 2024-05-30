package linkedlist

func (ll *linkedList[T]) clear() bool {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
	return true
}

// Clear removes all elements from the linked list and resets the head, tail, and size.
// It returns true if the operation was successful.
func (ll *LinkedList[T]) Clear() bool {
	return ll.linkedList_.clear()
}
