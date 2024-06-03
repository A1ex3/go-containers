package unorderedset

// Remove removes an element from the set and returns true if the element was removed, false if the element was not in the set.
func (s *UnorderedSet[T]) Remove(element T) bool {
	if !s.Contains(element) {
		return false
	}

	delete(s.Data, element)
	s.Size_--
	return true
}
