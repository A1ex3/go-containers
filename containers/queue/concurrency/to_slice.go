package concurrency

func (q *queueConcurrency[T]) toSlice() []T {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queueSync_.ToSlice()
}

func (q *QueueConcurrency[T]) ToSlice() []T {
	return q.queue_.toSlice()
}
