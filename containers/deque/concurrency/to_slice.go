package concurrency

func (d *dequeConcurrency[T]) toSlice() []T {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.dequeSync_.ToSlice()
}

func (d *DequeConcurrency[T]) ToSlice() []T {
	return d.deque_.toSlice()
}
