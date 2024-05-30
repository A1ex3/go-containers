package concurrency

func (ll *linkedList[T]) getByIndex(index int) (T, error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.GetByIndex(index)
}

// GetByIndex returns the data of the node at the specified index in the linked list.
// If the index is out of bounds, it returns a zero value of type T along with an error.
func (ll *LinkedList[T]) GetByIndex(index int) (T, error) {
	return ll.linkedList_.getByIndex(index)
}
