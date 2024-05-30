package linkedlist

import (
	"testing"
)

func TestInsertLast(t *testing.T) {
	var list ILinkedList[int] = NewLinkedList[int]()
	list.InsertLast(1)
	list.InsertLast(2)

	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}

	index, err := list.IndexOf(2)
	if err != nil || index != 1 {
		t.Errorf("Expected to find 2 at index 1, got index %d, err %v", index, err)
	}
}

func TestInsertFirst(t *testing.T) {
	var list ILinkedList[int] = NewLinkedList[int]()
	list.InsertFirst(1)
	list.InsertFirst(2)

	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}

	index, err := list.IndexOf(2)
	if err != nil || index != 0 {
		t.Errorf("Expected to find 2 at index 0, got index %d, err %v", index, err)
	}
}

func TestIndexOf(t *testing.T) {
	list := NewLinkedList[[]int]()
	list.InsertLast([]int{1, 2, 3})
	list.InsertLast([]int{4, 5, 6})

	index, err := list.IndexOf([]int{1, 2, 3})
	if err != nil || index != 0 {
		t.Errorf("Expected to find [1 2 3] at index 0, got index %d, err %v", index, err)
	}

	index, err = list.IndexOf([]int{4, 5, 6})
	if err != nil || index != 1 {
		t.Errorf("Expected to not find [4 5 6], got index %d, err %v", index, err)
	}
}

func TestRemove(t *testing.T) {
	var list ILinkedList[int] = NewLinkedList[int]()
	list.InsertLast(1)
	list.InsertLast(2)
	list.InsertLast(3)

	if !list.Remove(1) {
		t.Errorf("Expected to successfully remove element at index 1")
	}

	if list.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", list.Size())
	}

	index, err := list.IndexOf(2)
	if err == nil || index != -1 {
		t.Errorf("Expected to not find 2, got index %d, err %v", index, err)
	}
}

func TestClear(t *testing.T) {
	var list ILinkedList[int] = NewLinkedList[int]()
	list.InsertLast(1)
	list.InsertLast(2)
	list.InsertLast(3)

	if !list.Clear() {
		t.Errorf("Expected to successfully clear the list")
	}

	if list.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", list.Size())
	}
}

func TestContains(t *testing.T) {
	var list ILinkedList[int] = NewLinkedList[int]()
	list.InsertLast(1)
	list.InsertLast(2)

	if !list.Contains(1) {
		t.Errorf("Expected list to contain 1")
	}

	if list.Contains(3) {
		t.Errorf("Expected list to not contain 3")
	}
}

func TestSize(t *testing.T) {
	var list ILinkedList[int] = NewLinkedList[int]()
	if list.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", list.Size())
	}

	list.InsertLast(1)
	if list.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", list.Size())
	}
}
