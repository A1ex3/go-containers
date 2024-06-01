package concurrency

func (ll *linkedList[T]) swap(other *linkedList[T]) {
	ll.mu.Lock()
	other.mu.Lock()
	defer ll.mu.Unlock()
	defer other.mu.Unlock()

	ll.linkedListSync.Swap(other.linkedListSync)
}

func (ll *LinkedList[T]) Swap(other *LinkedList[T]) {
	ll.linkedList_.swap(other.linkedList_)
}
