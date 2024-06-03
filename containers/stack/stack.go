package stack

import (
	"github.com/a1ex3/go-containers/containers/deque"
)

// IStack defines the interface for a generic stack.
// It includes methods to push an element onto the stack,
// pop an element from the stack, and get the current size of the stack.
type IStack[T any] interface {
	Push(T)
	Pop() (T, error)
	Size() int
	Top() (T, error)
	ToSlice() []T
}

type StackDeque[T any] struct {
	Data *deque.Deque[T]
}

type StackSlice[T any] struct {
	Data []T
	Size_ int
}

// NewStackDeque initializes a new instance of Stack.
func NewStackDeque[T any]() *StackDeque[T] {
	return &StackDeque[T]{
		Data: deque.NewDeque[T](),
	}
}

func NewStackSlice[T any]() *StackSlice[T] {
	return &StackSlice[T]{
		Data: make([]T, 0),
		Size_: 0,
	}
}
