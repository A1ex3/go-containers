package concurrency

func (ll *linkedList[T]) removeFirst() bool {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.RemoveFirst()
}

func (ll *LinkedList[T]) RemoveFirst() bool {
	return ll.linkedList_.removeFirst()
}
