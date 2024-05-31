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
