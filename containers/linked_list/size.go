package linkedlist

func (ll *linkedList[T]) size_() int {
	return ll.size
}

// Size returns the number of elements in the linked list.
func (ll *LinkedList[T]) Size() int {
	return ll.linkedList_.size_()
}
