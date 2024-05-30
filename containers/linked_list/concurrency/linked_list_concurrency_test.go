package concurrency

import (
	"sync"
	"testing"

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
