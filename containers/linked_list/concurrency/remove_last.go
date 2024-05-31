package concurrency

func (ll *linkedList[T]) removeLast() bool {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.RemoveLast()
}

func (ll *LinkedList[T]) RemoveLast() bool {
	return ll.linkedList_.removeLast()
}
