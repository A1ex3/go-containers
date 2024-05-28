package concurrency

import (
	"fmt"
	"sync"
)

type stackConcurrency[T any] struct {
	mu   sync.Mutex // Mutex to ensure thread-safe access to the stack.
	data []T        // Slice to store the stack elements.
	size int        // Current size of the stack.
}

// StackConcurrency is a generic thread-safe stack.
// The type parameter T represents the type of elements stored in the stack.
type StackConcurrency[T any] struct {
	stack_ *stackConcurrency[T] // Internal stack implementation.
}

// Size returns the current number of elements in the stack.
// It locks the stack to ensure thread-safe access.
func (s *StackConcurrency[T]) Size() int {
	s.stack_.mu.Lock()
	defer s.stack_.mu.Unlock()
	return s.stack_.size
}

// newStackConcurrency initializes a new instance of stackConcurrency with the specified size.
// It returns a pointer to the newly created stackConcurrency.
func newStackConcurrency[T any](size int) *stackConcurrency[T] {
	return &stackConcurrency[T]{
		mu:   sync.Mutex{},
		data: make([]T, 0, size),
		size: 0,
	}
}

// NewStackConcurrency initializes a new instance of StackConcurrency with the specified size.
// If no size is provided, a default size of 10 is used.
// If more than one size argument is provided, it panics with an appropriate error message.
func NewStackConcurrency[T any](size ...int) *StackConcurrency[T] {
	var size_ int = 10

	if len(size) > 0 && len(size) < 2 {
		size_ = size[0]
	} else if len(size) > 1 {
		panic(fmt.Sprintf("Too many arguments, expected: 1, passed: %d", len(size)))
	}

	st := newStackConcurrency[T](size_)
	return &StackConcurrency[T]{
		stack_: st,
	}
}
