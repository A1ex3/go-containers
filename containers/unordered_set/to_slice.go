package unorderedset

// ToSlice converts the set to a slice.
func (s *UnorderedSet[T]) ToSlice() []T {
	keys := make([]T, 0, s.Size())
	for key := range s.Data {
		keys = append(keys, key.(T))
	}

	return keys
}
