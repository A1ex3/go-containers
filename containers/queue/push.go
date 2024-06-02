package queue

// Push adds an element to the back of the queue.
func (q *QueueDeque[T]) Push(value T) {
	q.data.PushLast(value)
}

func (s *QueueSlice[T]) Push(value T) {
	s.data = append(s.data, value)
	s.size++
}
