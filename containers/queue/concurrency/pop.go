package concurrency

func (q *QueueDequeConcurrency[T]) Pop() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queueSync_.Pop()
}

func (q *QueueSliceConcurrency[T]) Pop() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queueSync_.Pop()
}