package concurrency

import (
	"sync"

	linkedlist "github.com/a1ex3/go-containers/containers/linked_list"
)

type stackConcurrency[T any] struct {
	mu   sync.Mutex                // Mutex to ensure thread-safe access to the stack.
	data *linkedlist.LinkedList[T] // Slice to store the stack elements.
}

// StackConcurrency is a generic thread-safe stack.
// The type parameter T represents the type of elements stored in the stack.
type StackConcurrency[T any] struct {
	stack_ *stackConcurrency[T] // Internal stack implementation.
}

// newStackConcurrency initializes a new instance of stackConcurrency with the specified size.
// It returns a pointer to the newly created stackConcurrency.
func newStackConcurrency[T any]() *stackConcurrency[T] {
	linkedList := linkedlist.NewLinkedList[T]()
	return &stackConcurrency[T]{
		mu:   sync.Mutex{},
		data: linkedList,
	}
}

// NewStackConcurrency initializes a new instance of StackConcurrency with the specified size.
// If no size is provided, a default size of 10 is used.
// If more than one size argument is provided, it panics with an appropriate error message.
func NewStackConcurrency[T any]() *StackConcurrency[T] {
	st := newStackConcurrency[T]()
	return &StackConcurrency[T]{
		stack_: st,
	}
}
