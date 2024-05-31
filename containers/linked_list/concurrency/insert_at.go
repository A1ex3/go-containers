package concurrency

func (ll *linkedList[T]) insertAt(data T, index int) error {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.InsertAt(data, index)
}

func (ll *LinkedList[T]) InsertAt(data T, index int) error {
	return ll.linkedList_.insertAt(data, index)
}
