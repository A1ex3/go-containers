package queue

import (
	"testing"
)

// TestQueue_Push checks that items are correctly added to the queue and the queue size is incremented.
func TestQueue_Push(t *testing.T) {
	q := NewQueueDeque[int]()
	q.Push(10)
	q.Push(20)

	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}
}

// TestQueueue_Pop checks that items are correctly removed from the queue in the correct order and the queue size is reduced.
func TestQueue_Pop(t *testing.T) {
	q := NewQueueDeque[int]()
	q.Push(10)
	q.Push(20)

	val, err := q.Pop()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 10 {
		t.Errorf("expected value 10, got %d", val)
	}
	if q.Size() != 1 {
		t.Errorf("expected size 1, got %d", q.Size())
	}

	val, err = q.Pop()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("expected value 20, got %d", val)
	}
	if q.Size() != 0 {
		t.Errorf("expected size 0, got %d", q.Size())
	}
}

// TestQueueue_Size checks that the Size method returns the correct value in different queue states.
func TestQueue_Size(t *testing.T) {
	q := NewQueueDeque[int]()
	if q.Size() != 0 {
		t.Errorf("expected size 0, got %d", q.Size())
	}

	q.Push(10)
	if q.Size() != 1 {
		t.Errorf("expected size 1, got %d", q.Size())
	}

	q.Push(20)
	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}

	q.Pop()
	if q.Size() != 1 {
		t.Errorf("expected size 1, got %d", q.Size())
	}

	q.Pop()
	if q.Size() != 0 {
		t.Errorf("expected size 0, got %d", q.Size())
	}
}

// TestQueue_PopEmpty checks that an error is returned when trying to retrieve an item from an empty queue.
func TestQueue_PopEmpty(t *testing.T) {
	q := NewQueueDeque[int]()
	_, err := q.Pop()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestQueue(t *testing.T) {
	q := NewQueueDeque[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	front, err := q.Front()
	if err != nil {
		t.Errorf("Front() returned an error: %v", err)
	}
	if front != 1 {
		t.Errorf("Front() returned %d, expected 1", front)
	}

	popped, err := q.Pop()
	if err != nil {
		t.Errorf("Pop() returned an error: %v", err)
	}
	if popped != 1 {
		t.Errorf("Pop() returned %d, expected 1", popped)
	}

	size := q.Size()
	if size != 2 {
		t.Errorf("Size() returned %d, expected 2", size)
	}

	popped, err = q.Pop()
	if err != nil {
		t.Errorf("Pop() returned an error: %v", err)
	}
	if popped != 2 {
		t.Errorf("Pop() returned %d, expected 2", popped)
	}

	size = q.Size()
	if size != 1 {
		t.Errorf("Size() returned %d, expected 1", size)
	}

	popped, err = q.Pop()
	if err != nil {
		t.Errorf("Pop() returned an error: %v", err)
	}
	if popped != 3 {
		t.Errorf("Pop() returned %d, expected 3", popped)
	}

	size = q.Size()
	if size != 0 {
		t.Errorf("Size() returned %d, expected 0", size)
	}

	_, err = q.Pop()
	if err == nil {
		t.Error("Pop() did not return an error on empty queue")
	}
}

func TestQueueSlice_Push(t *testing.T) {
	q := NewQueueSlice[int]()
	q.Push(10)
	q.Push(20)

	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}
}

// TestQueueue_Pop checks that items are correctly removed from the queue in the correct order and the queue size is reduced.
func TestQueueSlice_Pop(t *testing.T) {
	q := NewQueueSlice[int]()
	q.Push(10)
	q.Push(20)

	val, err := q.Pop()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 10 {
		t.Errorf("expected value 10, got %d", val)
	}
	if q.Size() != 1 {
		t.Errorf("expected size 1, got %d", q.Size())
	}

	val, err = q.Pop()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("expected value 20, got %d", val)
	}
	if q.Size() != 0 {
		t.Errorf("expected size 0, got %d", q.Size())
	}
}

// TestQueueue_Size checks that the Size method returns the correct value in different queue states.
func TestQueueSlice_Size(t *testing.T) {
	q := NewQueueSlice[int]()
	if q.Size() != 0 {
		t.Errorf("expected size 0, got %d", q.Size())
	}

	q.Push(10)
	if q.Size() != 1 {
		t.Errorf("expected size 1, got %d", q.Size())
	}

	q.Push(20)
	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}

	q.Pop()
	if q.Size() != 1 {
		t.Errorf("expected size 1, got %d", q.Size())
	}

	q.Pop()
	if q.Size() != 0 {
		t.Errorf("expected size 0, got %d", q.Size())
	}
}

// TestQueue_PopEmpty checks that an error is returned when trying to retrieve an item from an empty queue.
func TestQueueSlice_PopEmpty(t *testing.T) {
	q := NewQueueSlice[int]()
	_, err := q.Pop()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestQueueSlice(t *testing.T) {
	q := NewQueueSlice[int]()

	q.Push(1)
	q.Push(2)
	q.Push(3)

	front, err := q.Front()
	if err != nil {
		t.Errorf("Front() returned an error: %v", err)
	}
	if front != 1 {
		t.Errorf("Front() returned %d, expected 1", front)
	}

	popped, err := q.Pop()
	if err != nil {
		t.Errorf("Pop() returned an error: %v", err)
	}
	if popped != 1 {
		t.Errorf("Pop() returned %d, expected 1", popped)
	}

	size := q.Size()
	if size != 2 {
		t.Errorf("Size() returned %d, expected 2", size)
	}

	popped, err = q.Pop()
	if err != nil {
		t.Errorf("Pop() returned an error: %v", err)
	}
	if popped != 2 {
		t.Errorf("Pop() returned %d, expected 2", popped)
	}

	size = q.Size()
	if size != 1 {
		t.Errorf("Size() returned %d, expected 1", size)
	}

	popped, err = q.Pop()
	if err != nil {
		t.Errorf("Pop() returned an error: %v", err)
	}
	if popped != 3 {
		t.Errorf("Pop() returned %d, expected 3", popped)
	}

	size = q.Size()
	if size != 0 {
		t.Errorf("Size() returned %d, expected 0", size)
	}

	_, err = q.Pop()
	if err == nil {
		t.Error("Pop() did not return an error on empty queue")
	}
}
