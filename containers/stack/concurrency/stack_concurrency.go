package concurrency

import (
	"sync"

	"github.com/a1ex3/go-containers/containers/stack"
)

type stackConcurrency[T any] struct {
	mu         sync.Mutex // Mutex to ensure thread-safe access to the stack.
	stackSync_ *stack.Stack[T]
}

// StackConcurrency is a generic thread-safe stack.
// The type parameter T represents the type of elements stored in the stack.
type StackConcurrency[T any] struct {
	stack_ *stackConcurrency[T]
}

// newStackConcurrency initializes a new instance of stackConcurrency with the specified size.
// It returns a pointer to the newly created stackConcurrency.
func newStackConcurrency[T any]() *stackConcurrency[T] {
	stack := stack.NewStack[T]()
	return &stackConcurrency[T]{
		mu:         sync.Mutex{},
		stackSync_: stack,
	}
}

// NewStackConcurrency initializes a new instance of StackConcurrency.
func NewStackConcurrency[T any]() *StackConcurrency[T] {
	st := newStackConcurrency[T]()
	return &StackConcurrency[T]{
		stack_: st,
	}
}
