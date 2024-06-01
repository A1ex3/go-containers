package linkedlist

func (ll *linkedList[T]) mergeBegin(other *linkedList[T]) {
	if other.head == nil {
		return
	}
	if ll.head == nil {
		ll.head = other.head
		ll.tail = other.tail
	} else {
		other.tail.next = ll.head
		ll.head = other.head
	}
	ll.size += other.size
}

func (ll *LinkedList[T]) MergeBegin(other *LinkedList[T]) {
	ll.linkedList_.mergeBegin(other.linkedList_)
}
