package linkedlist

func (ll *linkedList[T]) get() *LinkedListNode[T] {
	return ll.head
}

// Get returns the head node of the linked list.
func (ll *LinkedList[T]) Get() *LinkedListNode[T] {
	return ll.linkedList_.get()
}
