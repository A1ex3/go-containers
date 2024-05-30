package concurrency

func (ll *linkedList[T]) clear() bool {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.Clear()
}

// Clear removes all elements from the linked list and resets the head, tail, and size.
// It returns true if the operation was successful.
func (ll *LinkedList[T]) Clear() bool {
	return ll.linkedList_.clear()
}
