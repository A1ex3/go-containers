package linkedlist

func (ll *linkedList[T]) get() *linkedListNode[T] {
	return ll.head
}

// Get returns the head node of the linked list.
func (ll *LinkedList[T]) Get() *linkedListNode[T] {
	return ll.linkedList_.get()
}
