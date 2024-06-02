package queue

func (q *QueueDeque[T]) ToSlice() []T {
	return q.data.ToSlice()
}

func (q *QueueSlice[T]) ToSlice() []T {
	return q.data
}
