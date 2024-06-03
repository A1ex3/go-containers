package queue

func (q *QueueDeque[T]) ToSlice() []T {
	return q.Data.ToSlice()
}

func (q *QueueSlice[T]) ToSlice() []T {
	return q.Data
}
