package unorderedset

// Difference returns a new set containing elements present in the current set but absent in the specified set.
func (s *UnorderedSet[T]) Difference(other *UnorderedSet[T]) *UnorderedSet[T] {
	result := NewUnorderedSet[T]()
	for elem := range s.Data {
		if val := elem.(T); !other.Contains(val) {
			result.Add(val)
		}
	}
	return result
}
