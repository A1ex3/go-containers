package concurrency

import (
	"sync"

	linkedlist "github.com/a1ex3/go-containers/containers/linked_list"
)

// linkedList represents the actual implementation of the linked list.
type linkedList[T any] struct {
	mu         sync.Mutex
	linkedListSync *linkedlist.LinkedList[T]
}

// LinkedList wraps the linked list implementation to expose as a public API.
type LinkedList[T any] struct {
	linkedList_ *linkedList[T]
}

// newLinkedList creates a new instance of the linked list implementation.
func newLinkedList[T any]() *linkedList[T] {
	linkedListSync_ := linkedlist.NewLinkedList[T]()
	return &linkedList[T]{
		mu:         sync.Mutex{},
		linkedListSync: linkedListSync_,
	}
}

// NewLinkedList creates a new instance of the public LinkedList API.
func NewLinkedList[T any]() *LinkedList[T] {
	linkedList := newLinkedList[T]()

	return &LinkedList[T]{
		linkedList_: linkedList,
	}
}
