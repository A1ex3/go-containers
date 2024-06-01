package concurrency

func (ll *linkedList[T]) getFirst() (T, error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.GetFirst()
}

func (ll *LinkedList[T]) GetFirst() (T, error) {
	return ll.linkedList_.getFirst()
}
