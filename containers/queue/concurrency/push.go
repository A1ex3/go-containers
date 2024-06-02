package concurrency

func (q *QueueDequeConcurrency[T]) Push(value T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.queueSync_.Push(value)
}

func (q *QueueSliceConcurrency[T]) Push(value T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.queueSync_.Push(value)
}