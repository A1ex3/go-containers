package linkedlist

func (ll *linkedList[T]) insertFirst(data T) {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	newNode := &linkedListNode[T]{data: data}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
	ll.size++
}

// InsertFirst inserts a new element with the given data at the beginning of the linked list.
func (ll *LinkedList[T]) InsertFirst(data T) {
	ll.linkedList_.insertFirst(data)
}
