package queue

func (q *queue[T]) push(value T) {
	q.data.InsertLast(value)
}

// Push adds an element to the back of the queue.
func (q *Queue[T]) Push(value T) {
	q.queue_.push(value)
}
