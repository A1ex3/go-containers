package stack

import (
	linkedlist "github.com/a1ex3/go-containers/containers/linked_list"
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
	data *linkedlist.LinkedList[T] // Slice to store the stack elements.
}

type Stack[T any] struct {
	stack_ *stack[T]
}

func newStack[T any]() *stack[T] {
	linkedList := linkedlist.NewLinkedList[T]()
	return &stack[T]{
		data: linkedList,
	}
}

// NewStack initializes a new instance of Stack.
func NewStack[T any]() *Stack[T] {
	st := newStack[T]()
	return &Stack[T]{
		stack_: st,
	}
}
