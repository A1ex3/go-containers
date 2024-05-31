package concurrency

func (d *dequeConcurrency[T]) pushFirst(value T) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.dequeSync_.PushFirst(value)
}

func (d *DequeConcurrency[T]) PushFirst(value T) {
	d.deque_.pushFirst(value)
}
