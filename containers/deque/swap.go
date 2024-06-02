package deque

func (d *deque[T]) swap(other *Deque[T]) {
	d.data.Swap(other.deque_.data)
}

func (d *Deque[T]) Swap(other *Deque[T]) {
	d.deque_.swap(other)
}
