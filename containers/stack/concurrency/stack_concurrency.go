package concurrency

import (
	"sync"

	"github.com/a1ex3/go-containers/containers/stack"
)

// StackConcurrency is a generic thread-safe stack.
// The type parameter T represents the type of elements stored in the stack.
type StackDequeConcurrency[T any] struct {
	mu          sync.Mutex // Mutex to ensure thread-safe access to the stack.
	stackDeque_ *stack.StackDeque[T]
}

type StackSliceConcurrency[T any] struct {
	mu          sync.Mutex // Mutex to ensure thread-safe access to the stack.
	stackSlice_ *stack.StackSlice[T]
}

// NewStackDequeConcurrency initializes a new instance of StackConcurrency.
func NewStackDequeConcurrency[T any]() *StackDequeConcurrency[T] {
	return &StackDequeConcurrency[T]{
		stackDeque_: stack.NewStackDeque[T](),
	}
}

func NewStackSliceConcurrency[T any]() *StackSliceConcurrency[T] {
	return &StackSliceConcurrency[T]{
		stackSlice_: stack.NewStackSlice[T](),
	}
}
