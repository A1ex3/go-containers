package concurrency

func (ll *linkedList[T]) getLast() (T, error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.GetLast()
}

func (ll *LinkedList[T]) GetLast() (T, error) {
	return ll.linkedList_.getLast()
}
