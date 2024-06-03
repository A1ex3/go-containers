package unorderedset

// SymmetricDifference returns a new set containing elements present in one of the sets, but not both at the same time.
func (s *UnorderedSet[T]) SymmetricDifference(other *UnorderedSet[T]) *UnorderedSet[T] {
	result := NewUnorderedSet[T]()
	for elem := range s.Data {
		if val := elem.(T); !other.Contains(val) {
			result.Add(val)
		}
	}
	for elem := range other.Data {
		if val := elem.(T); !s.Contains(val) {
			result.Add(val)
		}
	}
	return result
}
