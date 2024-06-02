package queue

import "errors"

func (q *QueueDeque[T]) Front() (T, error) {
	lastElem, err := q.data.GetFirst()
	if err != nil {
		return lastElem, errors.New("queue is empty")
	} else {
		return lastElem, nil
	}
}

func (q *QueueSlice[T]) Front() (T, error) {
	if q.size == 0 {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return q.data[0], nil
}
