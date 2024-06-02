package queue

// Size returns the number of elements in the queue.
func (q *QueueDeque[T]) Size() int {
	return q.data.Size()
}

func (q *QueueSlice[T]) Size() int {
	return q.size
}
