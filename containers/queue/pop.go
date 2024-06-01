package queue

import (
	"errors"
)

func (q *queue[T]) pop() (T, error) {
	var zero T

	// Get the element at the front of the queue.
	res, err := q.data.GetFirst()
	if err != nil {
		return zero, errors.New("empty queue")
	}

	// Remove the element from the queue.
	if !q.data.RemoveFirst() {
		return zero, errors.New("failed to remove an element from the queue")
	}
	return res, nil
}

// Pop removes and returns the element at the front of the queue.
func (q *Queue[T]) Pop() (T, error) {
	return q.queue_.pop()
}
