package concurrency

func (ll *linkedList[T]) toSlice() []T {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.ToSlice()
}

func (ll *LinkedList[T]) ToSlice() []T {
	return ll.linkedList_.toSlice()
}
