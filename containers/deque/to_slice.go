package deque

func (d *deque[T]) toSlice() []T {
	return d.data.ToSlice()
}

func (d *Deque[T]) ToSlice() []T {
	return d.deque_.toSlice()
}
