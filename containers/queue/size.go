package queue

// Size returns the number of elements in the queue.
func (q *QueueDeque[T]) Size() int {
	return q.Data.Size()
}

func (q *QueueSlice[T]) Size() int {
	return q.Size_
}
