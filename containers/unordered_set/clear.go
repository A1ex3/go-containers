package unorderedset

// Clear clears the set.
func (s *UnorderedSet[T]) Clear() {
	s.Data = make(map[any]struct{})
	s.Size_ = 0
}
