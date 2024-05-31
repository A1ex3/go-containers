package concurrency

func (d *dequeConcurrency[T]) size_() int {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.dequeSync_.Size()
}

func (d *DequeConcurrency[T]) Size() int {
	return d.deque_.size_()
}
