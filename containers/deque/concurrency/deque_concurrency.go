package concurrency

import (
	"sync"

	"github.com/a1ex3/go-containers/containers/deque"
)

type dequeConcurrency[T any] struct {
	mu         sync.Mutex
	dequeSync_ *deque.Deque[T]
}

type DequeConcurrency[T any] struct {
	deque_ *dequeConcurrency[T]
}

func newDequeConcurrency[T any]() *dequeConcurrency[T] {
	deque := deque.NewDeque[T]()
	return &dequeConcurrency[T]{
		mu:         sync.Mutex{},
		dequeSync_: deque,
	}
}

func NewDequeConcurrency[T any]() *DequeConcurrency[T] {
	qe := newDequeConcurrency[T]()
	return &DequeConcurrency[T]{
		deque_: qe,
	}
}
