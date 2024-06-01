package linkedlist

import (
	"reflect"
	"testing"

	"github.com/a1ex3/go-containers/iterator"
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

func TestInsertAt(t *testing.T) {
	var ll ILinkedList[int] = NewLinkedList[int]()

	// Test inserting at index 0
	err := ll.InsertAt(5, 0)
	if err != nil {
		t.Errorf("Error inserting at index 0: %v", err)
	}
	if ll.Size() != 1 {
		t.Errorf("Size after inserting at index 0 should be 1, got %d", ll.Size())
	}

	// Test inserting at index out of bounds
	err = ll.InsertAt(10, 5)
	if err == nil {
		t.Error("Expected error when inserting at index out of bounds, but got nil")
	}

	// Test inserting at index within bounds
	err = ll.InsertAt(8, 1)
	if err != nil {
		t.Errorf("Error inserting at index 1: %v", err)
	}
	if ll.Size() != 2 {
		t.Errorf("Size after inserting at index 1 should be 2, got %d", ll.Size())
	}
}

func TestRemoveFirst(t *testing.T) {
	var ll ILinkedList[int] = NewLinkedList[int]()

	// Test removing from an empty list
	removed := ll.RemoveFirst()
	if removed {
		t.Error("Expected false for RemoveFirst on an empty list, but got true")
	}

	// Add elements to the list
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	// Test removing the first element
	removed = ll.RemoveFirst()
	if !removed {
		t.Error("Expected true for RemoveFirst on a non-empty list, but got false")
	}
	if ll.Size() != 2 {
		t.Errorf("Size after removing first element should be 2, got %d", ll.Size())
	}
}

func TestRemoveLast(t *testing.T) {
	var ll ILinkedList[int] = NewLinkedList[int]()

	// Test removing from an empty list
	removed := ll.RemoveLast()
	if removed {
		t.Error("Expected false for RemoveLast on an empty list, but got true")
	}

	// Add elements to the list
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	// Test removing the last element
	removed = ll.RemoveLast()
	if !removed {
		t.Error("Expected true for RemoveLast on a non-empty list, but got false")
	}
	if ll.Size() != 2 {
		t.Errorf("Size after removing last element should be 2, got %d", ll.Size())
	}
}

func TestReverse(t *testing.T) {
	var ll ILinkedList[int] = NewLinkedList[int]()

	// Test reversing an empty list
	ll.Reverse()

	// Add elements to the list
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	// Reverse the list
	ll.Reverse()

	// Check if the list is reversed correctly
	slice := ll.ToSlice()
	expected := []int{3, 2, 1}
	for i, val := range slice {
		if val != expected[i] {
			t.Errorf("Expected %d at index %d after reversing, got %d", expected[i], i, val)
		}
	}
}

func TestToSlice(t *testing.T) {
	var ll ILinkedList[int] = NewLinkedList[int]()

	// Test converting an empty list to slice
	slice := ll.ToSlice()
	if len(slice) != 0 {
		t.Errorf("Expected empty slice, got %v", slice)
	}

	// Add elements to the list
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	// Test converting non-empty list to slice
	slice = ll.ToSlice()
	expected := []int{1, 2, 3}
	for i, val := range slice {
		if val != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, val)
		}
	}
}

func TestForEach(t *testing.T) {
	var ll ILinkedList[int] = NewLinkedList[int]()

	// Add elements to the list
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	// Test applying a function to each element
	sum := 0
	ll.ForEach(func(val int) {
		sum += val
	})
	expectedSum := 6 // 1 + 2 + 3
	if sum != expectedSum {
		t.Errorf("Expected sum of %d, got %d", expectedSum, sum)
	}
}

func TestFind(t *testing.T) {
	var ll ILinkedList[int] = NewLinkedList[int]()

	// Add elements to the list
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	// Test finding an element
	found, err := ll.Find(func(val int) bool {
		return val == 2
	})
	if err != nil {
		t.Errorf("Error finding element: %v", err)
	}
	if found != 2 {
		t.Errorf("Expected to find element 2, got %d", found)
	}

	// Test not finding an element
	_, err = ll.Find(func(val int) bool {
		return val == 5
	})
	if err == nil {
		t.Error("Expected error when element not found, but got nil")
	}
}

func TestIterator(t *testing.T) {
	ll := NewLinkedList[int]()

	// Add elements to the list
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	// Get an iterator for the list
	var it iterator.IIterator[int] = NewIterator(ll)

	// Test hasNext and next
	expected := []int{1, 2, 3}
	for _, val := range expected {
		if !it.HasNext() {
			t.Error("Expected iterator to have next element, but got false")
		}
		nextVal, err := it.Next()
		if err != nil {
			t.Errorf("Error getting next element from iterator: %v", err)
		}
		if nextVal != val {
			t.Errorf("Expected next element to be %d, got %d", val, nextVal)
		}
	}

	// Test hasNext after iterating through all elements
	if it.HasNext() {
		t.Error("Expected iterator to not have next element after iterating through all elements, but got true")
	}
}

func TestUpdate(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedList[int]()

	// Insert some elements into the linked list
	ll.InsertLast(10)
	ll.InsertLast(20)
	ll.InsertLast(30)

	// Test updating data at valid index
	err := ll.Update(25, 1)
	if err != nil {
		t.Errorf("Failed to update data at valid index: %v", err)
	}

	// Test updating data at out-of-bounds index
	err = ll.Update(40, 3)
	if err == nil {
		t.Errorf("Updating data at out-of-bounds index should return an error")
	}

	// Test updating data at negative index
	err = ll.Update(50, -1)
	if err == nil {
		t.Errorf("Updating data at negative index should return an error")
	}

	// Test updating data at index 0
	err = ll.Update(5, 0)
	if err != nil {
		t.Errorf("Failed to update data at index 0: %v", err)
	}

	// Verify the updated data
	data, err := ll.GetByIndex(0)
	if err != nil {
		t.Errorf("Failed to get updated data: %v", err)
	}
	if data != 5 {
		t.Errorf("Updated data does not match")
	}
}

func TestSwap(t *testing.T) {
	ll1 := NewLinkedList[int]()
	ll2 := NewLinkedList[int]()

	ll1.InsertLast(1)
	ll1.InsertLast(2)
	ll1.InsertLast(3)
	ll1.InsertLast(4)

	ll2.InsertLast(5)
	ll2.InsertLast(6)
	ll2.InsertLast(7)

	ll1.Swap(ll2)

	if ll1.Size() != 3 {
		t.Errorf("Swap error List 1, expected size: %v, output: %v", 3, ll1.Size())
	}

	if ll2.Size() != 4 {
		t.Errorf("Swap error List 2, expected size: %v, output: %v", 4, ll2.Size())
	}

	ll1Slice := ll1.ToSlice()
	for i := 0; i < 3; i++ {
		if val := i + 5; val != ll1Slice[i] {
			t.Errorf("Swap error List 1, expected: %v, output: %v", val, ll1Slice[i])
		}
	}

	ll2Slice := ll2.ToSlice()
	for i := 0; i < 3; i++ {
		if val := i + 1; val != ll2Slice[i] {
			t.Errorf("Swap error List 2, expected: %v, output: %v", val, ll2Slice[i])
		}
	}
}

func TestMergeEnd(t *testing.T) {
	// Create the first linked list
	ll1 := NewLinkedList[int]()
	ll1.InsertLast(1)
	ll1.InsertLast(2)
	ll1.InsertLast(3)

	// Create the second linked list
	ll2 := NewLinkedList[int]()
	ll2.InsertLast(4)
	ll2.InsertLast(5)

	// Merge ll2 to the end of ll1
	ll1.MergeEnd(ll2)

	// Verify the merged list
	expected := []int{1, 2, 3, 4, 5}
	result := ll1.ToSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeEnd failed. Expected: %v, got: %v", expected, result)
	}
}
func TestMergeBegin(t *testing.T) {
	// Create the first linked list
	ll1 := NewLinkedList[int]()
	ll1.InsertLast(1)
	ll1.InsertLast(2)
	ll1.InsertLast(3)

	// Create the second linked list
	ll2 := NewLinkedList[int]()
	ll2.InsertLast(4)
	ll2.InsertLast(5)

	// Merge ll2 to the beginning of ll1
	ll1.MergeBegin(ll2)

	// Verify the merged list
	expected := []int{4, 5, 1, 2, 3}
	result := ll1.ToSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeBegin failed. Expected: %v, got: %v", expected, result)
	}
}

func TestGetFirst(t *testing.T) {
	list := NewLinkedList[int]()

	// Test when the list is empty
	_, err := list.GetFirst()
	if err == nil {
		t.Error("Expected error for empty list, got nil")
	}

	// Insert elements into the list
	list.InsertLast(1)
	list.InsertLast(2)
	list.InsertLast(3)

	// Test when the list has elements
	first, err := list.GetFirst()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if first != 1 {
		t.Errorf("Expected first element to be 1, got %d", first)
	}
}

func TestGetLast(t *testing.T) {
	list := NewLinkedList[int]()

	// Test when the list is empty
	_, err := list.GetLast()
	if err == nil {
		t.Error("Expected error for empty list, got nil")
	}

	// Insert elements into the list
	list.InsertLast(1)
	list.InsertLast(2)
	list.InsertLast(3)

	// Test when the list has elements
	last, err := list.GetLast()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if last != 3 {
		t.Errorf("Expected last element to be 3, got %d", last)
	}
}
