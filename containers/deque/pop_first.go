package deque

import (
	"errors"
)

func (d *deque[T]) popFirst() (T, error) {
	// Retrieve the first element
	res, err := d.data.GetFirst()
	if err != nil {
		return res, errors.New("empty deque")
	}

	// Remove the first element from the deque
	if !d.data.RemoveFirst() {
		return res, errors.New("failed to remove an element from the deque")
	}
	return res, nil // Return the removed element
}

// PopFirst removes and returns the first element of the deque.
func (d *Deque[T]) PopFirst() (T, error) {
	return d.deque_.popFirst()
}
