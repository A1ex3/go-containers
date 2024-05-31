package concurrency

func (ll *linkedList[T]) reverse() {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	ll.linkedListSync.Reverse()
}

func (ll *LinkedList[T]) Reverse() {
	ll.linkedList_.reverse()
}
