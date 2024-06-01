package concurrency

func (d *dequeConcurrency[T]) getLast() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.dequeSync_.GetLast()
}

func (d *DequeConcurrency[T]) GetLast() (T, error) {
	return d.deque_.getLast()
}
