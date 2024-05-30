package linkedlist

func (ll *linkedList[T]) insertLast(data T) {
	newNode := &LinkedListNode[T]{data: data}
	if ll.tail == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		ll.tail = newNode
	}
	ll.size++
}

// InsertLast inserts a new element with the given data at the end of the linked list.
func (ll *LinkedList[T]) InsertLast(data T) {
	ll.linkedList_.insertLast(data)
}
