package concurrency

func (ll *linkedList[T]) insertLast(data T) {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	ll.linkedListSync.InsertLast(data)
}

// InsertLast inserts a new element with the given data at the end of the linked list.
func (ll *LinkedList[T]) InsertLast(data T) {
	ll.linkedList_.insertLast(data)
}
