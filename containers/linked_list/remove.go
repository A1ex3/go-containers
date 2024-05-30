package linkedlist

func (ll *linkedList[T]) remove(index int) bool {
	if index < 0 || index >= ll.size {
		return false
	}
	if index == 0 {
		ll.head = ll.head.next
		if ll.size == 1 {
			ll.tail = nil
		}
	} else {
		prev := ll.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}
		prev.next = prev.next.next
		if index == ll.size-1 {
			ll.tail = prev
		}
	}
	ll.size--
	return true
}

// Remove removes the element at the specified index from the linked list.
// It returns true if the removal was successful, otherwise false.
func (ll *LinkedList[T]) Remove(index int) bool {
	return ll.linkedList_.remove(index)
}
