package queue

func (q *queue[T]) size_() int {
	return q.data.Size()
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return q.queue_.size_()
}
