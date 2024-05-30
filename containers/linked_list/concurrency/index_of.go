package concurrency

func (ll *linkedList[T]) indexOf(value T) (int, error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.IndexOf(value)
}

// IndexOf returns the index of the first occurrence of the specified value in the linked list.
// If the value is not found, it returns -1 along with an error.
func (ll *LinkedList[T]) IndexOf(value T) (int, error) {
	return ll.linkedList_.indexOf(value)
}
