package linkedlist

func (ll *linkedList[T]) toSlice() []T {
	slice := make([]T, 0, ll.size_())
	for current := ll.head; current != nil; current = current.next {
		slice = append(slice, current.data)
	}
	return slice
}

func (ll *LinkedList[T]) ToSlice() []T {
	return ll.linkedList_.toSlice()
}
