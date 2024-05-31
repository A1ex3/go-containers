package concurrency

func (ll *linkedList[T]) forEach(f func(T)) {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	ll.linkedListSync.ForEach(f)
}

func (ll *LinkedList[T]) ForEach(f func(T)) {
	ll.linkedList_.forEach(f)
}
