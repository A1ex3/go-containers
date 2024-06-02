package deque

import (
	linkedlist "github.com/a1ex3/go-containers/containers/linked_list"
)

// IDeque - interface for a double-ended queue (deque), defining methods for adding and removing elements from both ends.
type IDeque[T any] interface {
	PushFirst(T)          // Adds an element to the beginning of the deque.
	PushLast(T)           // Adds an element to the end of the deque.
	PopFirst() (T, error) // Removes and returns an element from the beginning of the deque. Returns an error if the deque is empty.
	PopLast() (T, error)  // Removes and returns an element from the end of the deque. Returns an error if the deque is empty.
	Size() int            // Returns the number of elements in the deque.
	GetFirst() (T, error)
	GetLast() (T, error)
	ToSlice() []T
}

type deque[T any] struct {
	data *linkedlist.LinkedList[T]
}

type Deque[T any] struct {
	deque_ *deque[T]
}

// newDeque creates a new instance of deque.
func newDeque[T any]() *deque[T] {
	linkedList := linkedlist.NewLinkedList[T]()
	return &deque[T]{
		data: linkedList,
	}
}

// NewDeque creates a new instance of Deque.
func NewDeque[T any]() *Deque[T] {
	dq := newDeque[T]()
	return &Deque[T]{
		deque_: dq,
	}
}
