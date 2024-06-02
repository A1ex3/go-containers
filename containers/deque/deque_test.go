package deque

import "testing"

func TestPushFirst(t *testing.T) {
	var q IDeque[int] = NewDeque[int]()
	q.PushFirst(1)
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}
	q.PushFirst(2)
	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}
}

func TestPushLast(t *testing.T) {
	var q IDeque[int] = NewDeque[int]()
	q.PushLast(1)
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}
	q.PushLast(2)
	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}
}

func TestPopFirst(t *testing.T) {
	var q IDeque[int] = NewDeque[int]()
	q.PushFirst(1)
	q.PushFirst(2)
	value, err := q.PopFirst()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 2 {
		t.Errorf("Expected value 2, got %d", value)
	}
	value, err = q.PopFirst()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 1 {
		t.Errorf("Expected value 1, got %d", value)
	}
	_, err = q.PopFirst()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestPopLast(t *testing.T) {
	var q IDeque[int] = NewDeque[int]()
	q.PushLast(1)
	q.PushLast(2)
	value, err := q.PopLast()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 2 {
		t.Errorf("Expected value 2, got %d", value)
	}
	value, err = q.PopLast()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 1 {
		t.Errorf("Expected value 1, got %d", value)
	}
	_, err = q.PopLast()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestSize(t *testing.T) {
	var q IDeque[int] = NewDeque[int]()
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
	q.PushFirst(1)
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}
	q.PushLast(2)
	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}
	q.PopFirst()
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}
	q.PopLast()
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
}

func TestGetFirst(t *testing.T) {
	deque := NewDeque[int]()

	// Test getting first element from an empty deque
	_, err := deque.GetFirst()
	if err == nil {
		t.Error("Expected error when getting first element from an empty deque, but got none")
	}

	// Test getting first element after pushing elements
	deque.PushLast(1)
	deque.PushLast(2)
	firstElement, err := deque.GetFirst()
	if err != nil {
		t.Errorf("Error getting first element from deque: %v", err)
	}
	if firstElement != 1 {
		t.Errorf("Expected first element to be 1, but got %d", firstElement)
	}
}

func TestGetLast(t *testing.T) {
	deque := NewDeque[int]()

	// Test getting last element from an empty deque
	_, err := deque.GetLast()
	if err == nil {
		t.Error("Expected error when getting last element from an empty deque, but got none")
	}

	// Test getting last element after pushing elements
	deque.PushLast(1)
	deque.PushLast(2)
	lastElement, err := deque.GetLast()
	if err != nil {
		t.Errorf("Error getting last element from deque: %v", err)
	}
	if lastElement != 2 {
		t.Errorf("Expected last element to be 2, but got %d", lastElement)
	}
}

func TestSwap(t *testing.T) {
	deq1 := NewDeque[int]()
	deq2 := NewDeque[int]()

	deq1.PushLast(1)
	deq1.PushLast(2)
	deq1.PushLast(3)
	deq1.PushLast(4)

	deq2.PushLast(5)
	deq2.PushLast(6)
	deq2.PushLast(7)

	deq1.Swap(deq2)

	if deq1.Size() != 3 {
		t.Errorf("Swap error List 1, expected size: %v, output: %v", 3, deq1.Size())
	}

	if deq2.Size() != 4 {
		t.Errorf("Swap error List 2, expected size: %v, output: %v", 4, deq2.Size())
	}

	deq1Slice := deq1.ToSlice()
	for i := 0; i < 3; i++ {
		if val := i + 5; val != deq1Slice[i] {
			t.Errorf("Swap error List 1, expected: %v, output: %v", val, deq1Slice[i])
		}
	}

	deq2Slice := deq2.ToSlice()
	for i := 0; i < 3; i++ {
		if val := i + 1; val != deq2Slice[i] {
			t.Errorf("Swap error List 2, expected: %v, output: %v", val, deq2Slice[i])
		}
	}
}
