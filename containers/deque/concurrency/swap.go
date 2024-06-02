package concurrency

func (d *dequeConcurrency[T]) swap(other *DequeConcurrency[T]) {
	d.mu.Lock()
	other.deque_.mu.Lock()
	defer d.mu.Unlock()
	defer other.deque_.mu.Unlock()

	d.dequeSync_.Swap(other.deque_.dequeSync_)
}

func (d *DequeConcurrency[T]) Swap(other *DequeConcurrency[T]) {
	d.deque_.swap(other)
}
