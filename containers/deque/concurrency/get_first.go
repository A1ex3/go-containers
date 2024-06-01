package concurrency

func (d *dequeConcurrency[T]) getFirst() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.dequeSync_.GetFirst()
}

func (d *DequeConcurrency[T]) GetFirst() (T, error) {
	return d.deque_.getFirst()
}
