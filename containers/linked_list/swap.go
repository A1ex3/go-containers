package linkedlist

func (ll *linkedList[T]) swap(other *linkedList[T]) {
	ll.head, other.head = other.head, ll.head
	ll.tail, other.tail = other.tail, ll.tail
	ll.size, other.size = other.size, ll.size
}

func (ll *LinkedList[T]) Swap(other *LinkedList[T]) {
	ll.linkedList_.swap(other.linkedList_)
}

