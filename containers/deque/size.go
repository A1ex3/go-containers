package deque

func (d *deque[T]) size_() int {
	return d.data.Size()
}

// Size returns the number of elements in the deque.
func (d *Deque[T]) Size() int {
	return d.deque_.size_()
}
