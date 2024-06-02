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
}

// queue is an implementation of the IQueue interface.
type queue[T any] struct {
	data *deque.Deque[T]
}

// Queue is a public wrapper around the private queue struct.
type Queue[T any] struct {
	queue_ *queue[T] // queue_ holds a pointer to the underlying private queue.
}

// newQueue creates a new instance of the queue struct.
func newQueue[T any]() *queue[T] {
	deque := deque.NewDeque[T]()
	return &queue[T]{
		data: deque,
	}
}

// NewQueue creates a new instance of the public Queue struct.
func NewQueue[T any]() *Queue[T] {
	qe := newQueue[T]() // Create a new queue.
	return &Queue[T]{
		queue_: qe, // Initialize the public Queue struct with the private queue.
	}
}
