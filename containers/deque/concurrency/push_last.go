package concurrency

func (d *dequeConcurrency[T]) pushLast(value T) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.dequeSync_.PushLast(value)
}

func (d *DequeConcurrency[T]) PushLast(value T) {
	d.deque_.pushLast(value)
}
