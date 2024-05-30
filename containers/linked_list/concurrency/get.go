package concurrency

import linkedlist "github.com/a1ex3/go-containers/containers/linked_list"

func (ll *linkedList[T]) get() *linkedlist.LinkedListNode[T] {
	ll.mu.Lock()
	defer ll.mu.Unlock()

	return ll.linkedListSync.Get()
}

// Get returns the head node of the linked list.
func (ll *LinkedList[T]) Get() *linkedlist.LinkedListNode[T] {
	return ll.linkedList_.get()
}
