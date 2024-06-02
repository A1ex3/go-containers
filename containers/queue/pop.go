package queue

import "errors"

// Pop removes and returns the element at the front of the queue.
func (q *QueueDeque[T]) Pop() (T, error) {
	return q.data.PopFirst()
}

func (q *QueueSlice[T]) Pop() (T, error) {
	if q.size == 0 {
		var zero T
		return zero, errors.New("queue is empty")
	}
	item := q.data[0]
	q.data = q.data[1:]
	q.size--
	return item, nil
}
