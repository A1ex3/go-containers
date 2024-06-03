package unorderedset

// Union returns a new set containing all elements from both sets.
func (s *UnorderedSet[T]) Union(other *UnorderedSet[T]) *UnorderedSet[T] {
	result := NewUnorderedSet[T]()
	for elem := range s.Data {
		result.Add(elem.(T))
	}
	for elem := range other.Data {
		result.Add(elem.(T))
	}
	return result
}
