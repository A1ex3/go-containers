package linkedlist

func (ll *linkedList[T]) forEach(f func(T)) {
	for current := ll.head; current != nil; current = current.next {
		f(current.data)
	}
}

func (ll *LinkedList[T]) ForEach(f func(T)) {
	ll.linkedList_.forEach(f)
}
