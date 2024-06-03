package queue

// Push adds an element to the back of the queue.
func (q *QueueDeque[T]) Push(value T) {
	q.Data.PushLast(value)
}

func (s *QueueSlice[T]) Push(value T) {
	s.Data = append(s.Data, value)
	s.Size_++
}
