package concurrency

func (ll *linkedList[T]) size_() int {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.Size()
}

// Size returns the number of elements in the linked list.
func (ll *LinkedList[T]) Size() int {
	return ll.linkedList_.size_()
}
