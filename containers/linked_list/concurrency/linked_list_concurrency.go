package linkedlist

import "sync"

// ILinkedList defines the interface for a generic linked list.
type ILinkedList[T any] interface {
	// InsertLast inserts data at the end of the linked list.
	InsertLast(data T)
	// InsertFirst inserts data at the beginning of the linked list.
	InsertFirst(data T)
	// IndexOf returns the index of the first occurrence of the specified value in the linked list.
	// Returns an error if the value is not found.
	IndexOf(value T) (int, error)
	// Get returns the head node of the linked list.
	Get() *linkedListNode[T]
	// GetByIndex returns the data at the specified index in the linked list.
	// Returns an error if the index is out of range.
	GetByIndex(index int) (T, error)
	// Remove removes the element at the specified index in the linked list.
	// Returns true if the element is removed successfully, false otherwise.
	Remove(index int) bool
	// Clear removes all elements from the linked list.
	// Returns true if the linked list is cleared successfully.
	Clear() bool
	// Contains checks whether the linked list contains the specified value.
	Contains(value T) bool
	// Size returns the number of elements in the linked list.
	Size() int
}

// linkedListNode represents a single node in the linked list.
type linkedListNode[T any] struct {
	data T
	next *linkedListNode[T]
}

// linkedList represents the actual implementation of the linked list.
type linkedList[T any] struct {
	mu   sync.Mutex
	head *linkedListNode[T]
	tail *linkedListNode[T]
	size int
}

// LinkedList wraps the linked list implementation to expose as a public API.
type LinkedList[T any] struct {
	linkedList_ *linkedList[T]
}

// newLinkedList creates a new instance of the linked list implementation.
func newLinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{
		mu:   sync.Mutex{},
		head: nil,
		tail: nil,
		size: 0,
	}
}

// NewLinkedList creates a new instance of the public LinkedList API.
func NewLinkedList[T any]() *LinkedList[T] {
	linkedList := newLinkedList[T]()

	return &LinkedList[T]{
		linkedList_: linkedList,
	}
}
