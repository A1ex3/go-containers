package concurrency

func (ll *linkedList[T]) find(predicate func(T) bool) (T, error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.Find(predicate)
}

func (ll *LinkedList[T]) Find(predicate func(T) bool) (T, error) {
	return ll.linkedList_.find(predicate)
}
