package concurrency

func (q *queueConcurrency[T]) front() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queueSync_.Front()
}

func (q *QueueConcurrency[T]) Front() (T, error) {
	return q.queue_.front()
}
