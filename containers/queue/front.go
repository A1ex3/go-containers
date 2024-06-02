package queue

import "errors"

func (q *queue[T]) front() (T, error) {
	lastElem, err := q.data.GetFirst()
	if err != nil {
		return lastElem, errors.New("queue is empty")
	} else {
		return lastElem, nil
	}
}

// Push adds an element to the back of the queue.
func (q *Queue[T]) Front() (T, error) {
	return q.queue_.front()
}
