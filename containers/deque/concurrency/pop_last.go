package concurrency

func (d *dequeConcurrency[T]) popLast() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.dequeSync_.PopLast()
}

func (d *DequeConcurrency[T]) PopLast() (T, error) {
	return d.deque_.popLast()
}
