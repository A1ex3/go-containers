package concurrency

func (d *dequeConcurrency[T]) popFirst() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.dequeSync_.PopFirst()
}

func (d *DequeConcurrency[T]) PopFirst() (T, error) {
	return d.deque_.popFirst()
}
