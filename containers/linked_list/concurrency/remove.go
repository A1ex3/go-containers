package concurrency

func (ll *linkedList[T]) remove(index int) bool {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.Remove(index)
}

// Remove removes the element at the specified index from the linked list.
// It returns true if the removal was successful, otherwise false.
func (ll *LinkedList[T]) Remove(index int) bool {
	return ll.linkedList_.remove(index)
}
