package deque

import (
	"errors"
)

func (d *deque[T]) getFirst() (T, error) {
	res, err := d.data.GetFirst()
	if err != nil {
		return res, errors.New("empty deque")
	}

	return res, nil
}

func (d *Deque[T]) GetFirst() (T, error) {
	return d.deque_.getFirst()
}
