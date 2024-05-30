package concurrency

func (q *queueConcurrency[T]) push(value T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.queueSync_.Push(value)
}

func (q *QueueConcurrency[T]) Push(value T) {
	q.queue_.push(value)
}
