package concurrency

import (
	"reflect"
	"sync"
	"testing"
	"time"

	linkedlist "github.com/a1ex3/go-containers/containers/linked_list"
)

func TestLinkedList_InsertLast(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		ll.InsertLast(1)
	}()

	go func() {
		defer wg.Done()
		ll.InsertLast(2)
	}()

	wg.Wait()

	if ll.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", ll.Size())
	}
}

func TestLinkedList_InsertFirst(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		ll.InsertFirst(1)
	}()

	go func() {
		defer wg.Done()
		ll.InsertFirst(2)
	}()

	wg.Wait()

	if ll.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", ll.Size())
	}
}

func TestLinkedList_GetByIndex(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	var result1 int
	go func() {
		defer wg.Done()
		var err error
		result1, err = ll.GetByIndex(0)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}()

	var result2 int
	go func() {
		defer wg.Done()
		var err error
		result2, err = ll.GetByIndex(1)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}()

	wg.Wait()

	if result1 != 1 {
		t.Errorf("Expected result1 to be 1, got %d", result1)
	}

	if result2 != 2 {
		t.Errorf("Expected result2 to be 2, got %d", result2)
	}
}

func TestLinkedList_Remove(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		ll.Remove(1)
	}()

	go func() {
		defer wg.Done()
		ll.Remove(0)
	}()

	wg.Wait()

	if ll.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", ll.Size())
	}

	// Test removal of non-existent index
	if ok := ll.Remove(10); ok {
		t.Error("Expected Remove(10) to return false")
	}
}

func TestLinkedList_Clear(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)

	ll.Clear()

	if ll.Size() != 0 {
		t.Errorf("Expected size to be 0 after clearing, got %d", ll.Size())
	}
}

func TestLinkedList_Contains(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)

	if !ll.Contains(1) {
		t.Error("Expected Contains(1) to return true")
	}

	if ll.Contains(3) {
		t.Error("Expected Contains(3) to return false")
	}
}

func TestLinkedList_Size(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)

	if ll.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", ll.Size())
	}
}

func TestInsertAt(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	err := ll.InsertAt(4, 1)
	if err != nil || ll.Size() != 4 || ll.ToSlice()[1] != 4 {
		t.Errorf("Expected insertion at index 1 to be successful")
	}

	err = ll.InsertAt(5, 5)
	if err == nil {
		t.Errorf("Expected out of bounds error for insertion")
	}
}

func TestRemoveFirst(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	removed := ll.RemoveFirst()
	if !removed || ll.Size() != 2 || ll.ToSlice()[0] != 2 {
		t.Errorf("Expected removal of first element to be successful")
	}

	ll.Clear()
	removed = ll.RemoveFirst()
	if removed {
		t.Errorf("Expected removal of first element to be unsuccessful on empty list")
	}
}

func TestRemoveLast(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	removed := ll.RemoveLast()
	if !removed || ll.Size() != 2 || ll.ToSlice()[1] != 2 {
		t.Errorf("Expected removal of last element to be successful")
	}

	ll.Clear()
	removed = ll.RemoveLast()
	if removed {
		t.Errorf("Expected removal of last element to be unsuccessful on empty list")
	}
}

func TestReverse(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)
	ll.Reverse()

	slice := ll.ToSlice()
	if slice[0] != 3 || slice[1] != 2 || slice[2] != 1 {
		t.Errorf("Expected reversed list to be [3, 2, 1], got %v", slice)
	}
}

func TestToSlice(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	slice := ll.ToSlice()
	if slice[0] != 1 || slice[1] != 2 || slice[2] != 3 {
		t.Errorf("Expected slice to be [1, 2, 3], got %v", slice)
	}
}

func TestForEach(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	var sum int
	ll.ForEach(func(val int) {
		sum += val
	})

	if sum != 6 {
		t.Errorf("Expected sum to be 6, got %d", sum)
	}
}

func TestFind(t *testing.T) {
	var ll linkedlist.ILinkedList[int] = NewLinkedListConcurrency[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	val, err := ll.Find(func(val int) bool {
		return val > 1
	})

	if val != 2 || err != nil {
		t.Errorf("Expected first value greater than 1 to be 2, got %v, error: %v", val, err)
	}

	_, err = ll.Find(func(val int) bool {
		return val > 3
	})

	if err == nil {
		t.Errorf("Expected error for predicate not satisfied")
	}
}

func TestLinkedListIterator(t *testing.T) {
	ll := linkedlist.NewLinkedList[int]()
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)

	iter := NewIteratorConcurrency[int](ll)

	var sum int
	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for iter.HasNext() {
				val, err := iter.Next()
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				sum += val
			}
		}()
	}

	wg.Wait()

	expectedSum := 1 + 2 + 3
	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, got %d", expectedSum, sum)
	}
}

func TestUpdate(t *testing.T) {
	// Create a new linked list
	ll := NewLinkedListConcurrency[int]()

	// Use a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Insert elements into the linked list concurrently
	for i := 0; i < 2; i++ { // Reduced to 2 goroutines to simulate a potential deadlock
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// Acquire the lock to insert an element
			ll.InsertLast(n)
		}(i)
	}

	// Wait for all insertion goroutines to finish
	wg.Wait()

	// Use another WaitGroup for updating data concurrently
	var updateWg sync.WaitGroup

	// Update data in the linked list concurrently
	for i := 0; i < 2; i++ { // Reduced to 2 goroutines to simulate a potential deadlock
		updateWg.Add(1)
		go func(n int) {
			defer updateWg.Done()
			// Acquire the lock to update data
			err := ll.Update(n*10, n)
			if err != nil {
				t.Errorf("Failed to update data at index %d: %v", n, err)
			}
		}(i)
	}

	// Wait for all update goroutines to finish
	updateWg.Wait()

	// Verify that the data is updated correctly
	for i := 0; i < 2; i++ {
		data, err := ll.GetByIndex(i)
		if err != nil {
			t.Errorf("Failed to get data at index %d: %v", i, err)
		}
		expected := i * 10
		if data != expected {
			t.Errorf("Data at index %d does not match expected value. Got: %d, Expected: %d", i, data, expected)
		}
	}
}

func TestSwap(t *testing.T) {
	var wg sync.WaitGroup

	ll1 := NewLinkedListConcurrency[int]()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ll1.InsertLast(n)
		}(i)
	}
	wg.Wait()

	ll2 := NewLinkedListConcurrency[int]()
	for i := 0; i < 4; i++ {
		ll2.InsertLast(i + 3)
	}

	t.Logf("Before swap: ll1 = %v, ll2 = %v", ll1.ToSlice(), ll2.ToSlice())

	wg.Add(2)

	go func() {
		defer wg.Done()
		ll1.Swap(ll2)
	}()

	go func() {
		defer wg.Done()
		ll2.Swap(ll1)
	}()

	c := make(chan struct{})

	go func() {
		wg.Wait()
		close(c)
	}()

	select {
	case <-c:
		t.Log("Swap completed without deadlock")
	case <-time.After(2 * time.Second):
		t.Fatal("Potential deadlock detected")
	}

	t.Logf("After swap: ll1 = %v, ll2 = %v", ll1.ToSlice(), ll2.ToSlice())
}

func TestMergeEnd(t *testing.T) {
	// Create the first linked list
	ll1 := NewLinkedListConcurrency[int]()
	ll1.InsertLast(1)
	ll1.InsertLast(2)
	ll1.InsertLast(3)

	// Create the second linked list
	ll2 := NewLinkedListConcurrency[int]()
	ll2.InsertLast(4)
	ll2.InsertLast(5)

	// Create a wait group
	var wg sync.WaitGroup
	wg.Add(1)

	// Merge ll2 to the end of ll1 using goroutine
	go func() {
		ll1.MergeEnd(ll2)
		wg.Done()
	}()

	// Wait for the goroutine to finish
	wg.Wait()

	// Verify the merged list
	expected := []int{1, 2, 3, 4, 5}
	result := ll1.ToSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeEnd failed. Expected: %v, got: %v", expected, result)
	}
}

func TestMergeBegin(t *testing.T) {
	// Create the first linked list
	ll1 := NewLinkedListConcurrency[int]()
	ll1.InsertLast(1)
	ll1.InsertLast(2)
	ll1.InsertLast(3)

	// Create the second linked list
	ll2 := NewLinkedListConcurrency[int]()
	ll2.InsertLast(4)
	ll2.InsertLast(5)

	// Create a wait group
	var wg sync.WaitGroup
	wg.Add(1)

	// Merge ll2 to the beginning of ll1 using goroutine
	go func() {
		ll1.MergeBegin(ll2)
		wg.Done()
	}()

	// Wait for the goroutine to finish
	wg.Wait()

	// Verify the merged list
	expected := []int{4, 5, 1, 2, 3}
	result := ll1.ToSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeBegin failed. Expected: %v, got: %v", expected, result)
	}
}
