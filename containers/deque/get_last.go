package deque

import (
	"errors"
)

func (d *deque[T]) getLast() (T, error) {
	var zero T

	res, err := d.data.GetLast()
	if err != nil {
		return zero, errors.New("empty deque")
	}

	return res, nil
}

func (d *Deque[T]) GetLast() (T, error) {
	return d.deque_.getLast()
}
