package concurrency

func (q *QueueDequeConcurrency[T]) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queueSync_.Size()
}

func (q *QueueSliceConcurrency[T]) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queueSync_.Size()
}
