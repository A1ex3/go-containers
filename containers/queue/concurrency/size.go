package concurrency

func (q *queueConcurrency[T]) size_() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queueSync_.Size()
}

func (q *QueueConcurrency[T]) Size() int {
	return q.queue_.size_()
}
