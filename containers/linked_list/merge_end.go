package linkedlist

func (ll *linkedList[T]) mergeEnd(other *linkedList[T]) {
	if ll.head == nil {
		ll.head = other.head
	} else {
		ll.tail.next = other.head
	}
	ll.tail = other.tail
	ll.size += other.size
}

func (ll *LinkedList[T]) MergeEnd(other *LinkedList[T]) {
	ll.linkedList_.mergeEnd(other.linkedList_)
}
