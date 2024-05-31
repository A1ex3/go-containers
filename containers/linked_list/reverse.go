package linkedlist

func (ll *linkedList[T]) reverse() {
	var prev, next *LinkedListNode[T]
	current := ll.head
	ll.tail = ll.head
	for current != nil {
		next = current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	ll.head = prev
}

func (ll *LinkedList[T]) Reverse() {
	ll.linkedList_.reverse()
}
