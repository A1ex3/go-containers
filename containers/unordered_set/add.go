package unorderedset

// Add adds an element to the set and returns true if the element was added, false if the element was already present.
func (s *UnorderedSet[T]) Add(element T) bool {
	if s.Contains(element) {
		return false
	}

	s.Data[element] = struct{}{}
	s.Size_++
	return true
}
