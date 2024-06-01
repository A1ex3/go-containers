package concurrency

func (ll *linkedList[T]) mergeBegin(other *linkedList[T]) {
	ll.mu.Lock()
	other.mu.Lock()
	defer ll.mu.Unlock()
	defer other.mu.Unlock()

	ll.linkedListSync.MergeBegin(other.linkedListSync)
}

func (ll *LinkedList[T]) MergeBegin(other *LinkedList[T]) {
	ll.linkedList_.mergeBegin(other.linkedList_)
}
