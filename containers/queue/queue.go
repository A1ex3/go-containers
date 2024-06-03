package queue

import (
	"github.com/a1ex3/go-containers/containers/deque"
)

// IQueue defines the interface for a generic queue.
type IQueue[T any] interface {
	Push(T)          // Push adds an element to the queue.
	Pop() (T, error) // Pop removes and returns the element at the front of the queue.
	Size() int       // Size returns the number of elements in the queue.
	Front() (T, error)
	ToSlice() []T
}

// Queue is a public wrapper around the private queue struct.
type QueueDeque[T any] struct {
	Data *deque.Deque[T]
}

type QueueSlice[T any] struct {
	Data []T
	Size_ int
}

// NewQueueDeque creates a new instance of the queue struct.
func NewQueueDeque[T any]() *QueueDeque[T] {
	return &QueueDeque[T]{
		Data: deque.NewDeque[T](),
	}
}

func NewQueueSlice[T any]() *QueueSlice[T] {
	return &QueueSlice[T]{
		Data: make([]T, 0),
		Size_: 0,
	}
}
