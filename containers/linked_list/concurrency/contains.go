package concurrency

func (ll *linkedList[T]) contains(value T) bool {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.Contains(value)
}

// Contains checks whether the specified value exists in the linked list.
// It iterates through each node and compares the node's data with the given value using reflect.DeepEqual.
// It returns true if the value is found, otherwise false.
func (ll *LinkedList[T]) Contains(value T) bool {
	return ll.linkedList_.contains(value)
}
