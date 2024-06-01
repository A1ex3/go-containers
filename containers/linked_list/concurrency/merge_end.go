package concurrency

func (ll *linkedList[T]) mergeEnd(other *linkedList[T]) {
	ll.mu.Lock()
	other.mu.Lock()
	defer ll.mu.Unlock()
	defer other.mu.Unlock()

	ll.linkedListSync.MergeEnd(other.linkedListSync)
}

func (ll *LinkedList[T]) MergeEnd(other *LinkedList[T]) {
	ll.linkedList_.mergeEnd(other.linkedList_)
}
