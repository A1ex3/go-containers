package concurrency

func (ll *linkedList[T]) update(data T, index int) error {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.Update(data, index)
}

func (ll *LinkedList[T]) Update(data T, index int) error {
	return ll.linkedList_.update(data, index)
}
