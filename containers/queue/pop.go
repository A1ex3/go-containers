package queue

import "errors"

// Pop removes and returns the element at the front of the queue.
func (q *QueueDeque[T]) Pop() (T, error) {
	return q.Data.PopFirst()
}

func (q *QueueSlice[T]) Pop() (T, error) {
	if q.Size_ == 0 {
		var zero T
		return zero, errors.New("queue is empty")
	}
	item := q.Data[0]
	q.Data = q.Data[1:]
	q.Size_--
	return item, nil
}
