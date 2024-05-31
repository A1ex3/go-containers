package deque

func (d *deque[T]) pushFirst(value T) {
	d.data.InsertFirst(value)
}

// PushFirst adds an element to the beginning of the deque.
func (d *Deque[T]) PushFirst(value T) {
	d.deque_.pushFirst(value)
}
