package deque

import (
	"errors"
)

func (d *deque[T]) popLast() (T, error) {
	var zero T // Default value of type T

	// Check if the deque is empty
	if d.data.Size() == 0 {
		return zero, errors.New("empty deque")
	}

	// Retrieve the last element
	res, err := d.data.GetByIndex(d.data.Size() - 1)
	if err != nil {
		return zero, err
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
