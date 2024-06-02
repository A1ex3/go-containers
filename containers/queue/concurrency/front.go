package concurrency

func (q *QueueDequeConcurrency[T]) Front() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queueSync_.Front()
}

func (q *QueueSliceConcurrency[T]) Front() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queueSync_.Front()
}