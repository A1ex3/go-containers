package concurrency

func (q *QueueDequeConcurrency[T]) ToSlice() []T {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queueSync_.ToSlice()
}

func (q *QueueSliceConcurrency[T]) ToSlice() []T {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queueSync_.ToSlice()
}
