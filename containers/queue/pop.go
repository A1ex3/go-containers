package queue

func (q *queue[T]) pop() (T, error) {
	return q.data.PopFirst()
}

// Pop removes and returns the element at the front of the queue.
func (q *Queue[T]) Pop() (T, error) {
	return q.queue_.pop()
}
