package concurrency

func (ll *linkedList[T]) insertFirst(data T) {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	ll.linkedListSync.InsertFirst(data)
}

// InsertFirst inserts a new element with the given data at the beginning of the linked list.
func (ll *LinkedList[T]) InsertFirst(data T) {
	ll.linkedList_.insertFirst(data)
}
