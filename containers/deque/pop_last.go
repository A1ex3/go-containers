package deque

import (
	"errors"
)

func (d *deque[T]) popLast() (T, error) {
	var zero T // Default value of type T

	// Retrieve the last element
	res, err := d.data.GetLast()
	if err != nil {
		return zero, errors.New("empty deque")
	}

	// Remove the last element from the deque
	if !d.data.RemoveLast() {
		return zero, errors.New("failed to remove an element from the deque")
	}
	return res, nil // Return the removed element
}

// PopLast removes and returns the last element of the deque.
func (d *Deque[T]) PopLast() (T, error) {
	return d.deque_.popLast()
}
