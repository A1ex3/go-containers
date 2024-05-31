package linkedlist

func (ll *linkedList[T]) removeLast() bool {
	return ll.remove(ll.size_() - 1)
}

func (ll *LinkedList[T]) RemoveLast() bool {
	return ll.linkedList_.removeLast()
}
