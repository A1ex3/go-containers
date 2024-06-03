package unorderedset

// Intersection returns a new set containing elements present in both sets.
func (s *UnorderedSet[T]) Intersection(other *UnorderedSet[T]) *UnorderedSet[T] {
	result := NewUnorderedSet[T]()
	for elem := range s.Data {
		if val := elem.(T); other.Contains(val) {
			result.Add(val)
		}
	}
	return result
}
