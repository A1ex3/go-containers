package stack

import (
	"fmt"
)

// IStack defines the interface for a generic stack.
// It includes methods to push an element onto the stack,
// pop an element from the stack, and get the current size of the stack.
type IStack[T any] interface {
	Push(T)
	Pop() (T, error)
	Size() int
}

// stack is a simple generic stack implementation.
// The type parameter T represents the type of elements stored in the stack.
type stack[T any] struct {
	data []T // Slice to store the stack elements.
	size int // Current size of the stack.
}

type Stack[T any] struct {
	stack_ *stack[T]
}

// Size returns the current number of elements in the stack.
func (s *Stack[T]) Size() int {
	return s.stack_.size
}

func newStack[T any](size int) *stack[T] {
	return &stack[T]{
		data: make([]T, 0, size),
		size: 0,
	}
}

// NewStack initializes a new instance of Stack with the specified size.
// If no size is provided, a default size of 10 is used.
// If more than one size argument is provided, it panics with an appropriate error message.
func NewStack[T any](size ...int) *Stack[T] {
	var size_ int = 10

	if len(size) > 0 && len(size) < 2 {
		size_ = size[0]
	} else if len(size) > 1 {
		panic(fmt.Sprintf("Too many arguments, expected: 1, passed: %d", len(size)))
	}

	st := newStack[T](size_)
	return &Stack[T]{
		stack_: st,
	}
}
