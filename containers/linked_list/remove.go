package linkedlist

func (ll *linkedList[T]) remove(index int) bool {
	if index < 0 || index >= ll.size_() {
		return false
	}

	var current *LinkedListNode[T]
	if index < ll.size_()/2 {
		current = ll.head
		for i := 0; i < index; i++ {
			current = current.next
		}
	} else {
		current = ll.tail
		for i := ll.size_() - 1; i > index; i-- {
			current = current.prev
		}
	}

	if current.prev != nil {
		current.prev.next = current.next
	} else {
		ll.head = current.next
	}

	if current.next != nil {
		current.next.prev = current.prev
	} else {
		ll.tail = current.prev
	}

	ll.size--
	return true
}

// Remove removes the element at the specified index from the linked list.
// It returns true if the removal was successful, otherwise false.
func (ll *LinkedList[T]) Remove(index int) bool {
	return ll.linkedList_.remove(index)
}
