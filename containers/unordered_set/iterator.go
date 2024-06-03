package unorderedset

import "errors"

type UnorderedSetIterator[T any] struct {
	set   *UnorderedSet[T]
	keys  []T
	index int
	size  int
}

// NewIterator returns a slice of elements of the set.
func (s *UnorderedSet[T]) NewIterator() *UnorderedSetIterator[T] {
	return &UnorderedSetIterator[T]{set: s, keys: s.ToSlice(), index: -1, size: s.Size()}
}

func (it *UnorderedSetIterator[T]) HasNext() bool {
	return it.index+1 < it.size
}

func (it *UnorderedSetIterator[T]) Next() (T, error) {
	if !it.HasNext() {
		var zero T
		return zero, errors.New("no more elements")
	}

	it.index++
	return it.keys[it.index], nil
}
