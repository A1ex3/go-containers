package unorderedset

// Contains checks for the presence of an element in the set.
func (s *UnorderedSet[T]) Contains(element T) bool {
	_, exists := s.Data[element]
	return exists
}
