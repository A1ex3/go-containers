package deque

func (d *deque[T]) pushLast(value T) {
	d.data.InsertLast(value)
}

// PushLast adds an element to the end of the deque.
func (d *Deque[T]) PushLast(value T) {
	d.deque_.pushLast(value)
}
