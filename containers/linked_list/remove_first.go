package linkedlist

func (ll *linkedList[T]) removeFirst() bool {
	return ll.remove(0)
}

func (ll *LinkedList[T]) RemoveFirst() bool {
	return ll.linkedList_.removeFirst()
}
