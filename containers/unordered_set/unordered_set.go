package unorderedset

// IUnorderedSet defines the interface for an unordered set.
type IUnorderedSet[T any] interface {
	Add(T) bool      // Add adds an element to the set.
	Remove(T) bool   // Remove removes an element from the set.
	Contains(T) bool // Contains checks if an element is present in the set.
	Size() int       // Size returns the number of elements in the set.
	Clear()          // Clear removes all elements from the set.
	ToSlice() []T    // ToSlice returns the elements of the set as a slice.
}

// UnorderedSet is an implementation of the IUnorderedSet interface.
type UnorderedSet[T any] struct {
	Data  map[any]struct{} // Data stores the elements of the set.
	Size_ int              // Size_ keeps track of the number of elements in the set.
}

// NewUnorderedSet creates and initializes a new UnorderedSet instance.
func NewUnorderedSet[T any]() *UnorderedSet[T] {
	return &UnorderedSet[T]{
		Data:  make(map[any]struct{}),
		Size_: 0,
	}
}
