package linkedlist

import "fmt"

func (ll *linkedList[T]) update(data T, index int) error {
	// Check if the index is out of bounds
	if index < 0 || index >= ll.size_() {
		return fmt.Errorf("index out of bounds")
	}

	// Traverse the list to find the node at the specified index
	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	// Update the data of the node
	current.data = data

	return nil
}

func (ll *LinkedList[T]) Update(data T, index int) error {
	return ll.linkedList_.update(data, index)
}
