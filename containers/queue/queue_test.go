package queue

import (
	"testing"
)

// TestQueue_Push checks that items are correctly added to the queue and the queue size is incremented.
func TestQueue_Push(t *testing.T) {
	q := NewQueue[int]()
	q.Push(10)
	q.Push(20)

	if q.Size() != 2 {
		t.Errorf("expected size 2, got %d", q.Size())
	}
}

// TestQueueue_Pop checks that items are correctly removed from the queue in the correct order and the queue size is reduced.
func TestQueue_Pop(t *testing.T) {
	q := NewQueue[int]()
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
	q := NewQueue[int]()
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
	q := NewQueue[int]()
	_, err := q.Pop()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
