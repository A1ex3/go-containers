package concurrency

func (q *queueConcurrency[T]) pop() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queueSync_.Pop()
}

func (q *QueueConcurrency[T]) Pop() (T, error) {
	return q.queue_.pop()
}
