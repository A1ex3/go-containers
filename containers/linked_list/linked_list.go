package linkedlist

// ILinkedList is an interface defining operations that can be performed on a linked list.
type ILinkedList[T any] interface {
	// InsertLast inserts a new element with the given data at the end of the linked list.
	InsertLast(data T)

	// InsertFirst inserts a new element with the given data at the beginning of the linked list.
	InsertFirst(data T)

	// IndexOf returns the index of the first occurrence of the specified value in the linked list.
	// If the value is not found, it returns -1 along with an error.
	IndexOf(value T) (int, error)

	// Get returns the head node of the linked list.
	Get() *LinkedListNode[T]

	// GetByIndex returns the data of the node at the specified index in the linked list.
	// If the index is out of bounds, it returns an error.
	GetByIndex(index int) (T, error)

	// Remove removes the element at the specified index from the linked list.
	// It returns true if the removal was successful, otherwise false.
	Remove(index int) bool

	// Clear removes all elements from the linked list.
	// It returns true if the operation was successful, otherwise false.
	Clear() bool

	// Contains checks if the specified value exists in the linked list.
	// It returns true if the value is found, otherwise false.
	Contains(value T) bool

	// Size returns the number of elements in the linked list.
	Size() int
}

// linkedListNode represents a node in the linked list containing generic data and a reference to the next node.
type LinkedListNode[T any] struct {
	data T
	next *LinkedListNode[T]
}

// linkedList represents the linked list structure with references to the head, tail, and size.
type linkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
	size int
}

// LinkedList wraps the linked list and provides a way to access its methods.
type LinkedList[T any] struct {
	linkedList_ *linkedList[T]
}

// newLinkedList creates and initializes a new linked list instance.
func newLinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// NewLinkedList creates and returns a new LinkedList instance.
func NewLinkedList[T any]() *LinkedList[T] {
	// Create a new linked list using newLinkedList function.
	linkedList := newLinkedList[T]()

	// Wrap the linked list in a LinkedList struct and return.
	return &LinkedList[T]{
		linkedList_: linkedList,
	}
}
