package queue

func (q *queue[T]) toSlice() []T {
	return q.data.ToSlice()
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) ToSlice() []T {
	return q.queue_.toSlice()
}
