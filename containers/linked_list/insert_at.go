package linkedlist

import "errors"

func (ll *linkedList[T]) insertAt(data T, index int) error {
	if index < 0 || index > ll.size_() {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		ll.insertFirst(data)
		return nil
	}
	if index == ll.size {
		ll.insertLast(data)
		return nil
	}
	newNode := &LinkedListNode[T]{data: data}
	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	newNode.next = current
	newNode.prev = current.prev
	current.prev.next = newNode
	current.prev = newNode
	ll.size++
	return nil
}

func (ll *LinkedList[T]) InsertAt(data T, index int) error {
	return ll.linkedList_.insertAt(data, index)
}
