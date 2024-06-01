package concurrency

import (
	"sync"
	"testing"
)

func TestConcurrentPushFirst(t *testing.T) {
	q := NewDequeConcurrency[int]()
	var wg sync.WaitGroup

	// PushFirst concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			q.PushFirst(n)
		}(i)
	}

	wg.Wait()

	if q.Size() != 1000 {
		t.Errorf("Expected size 1000, got %d", q.Size())
	}
}

func TestConcurrentPushLast(t *testing.T) {
	q := NewDequeConcurrency[int]()
	var wg sync.WaitGroup

	// PushLast concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			q.PushLast(n)
		}(i)
	}

	wg.Wait()

	if q.Size() != 1000 {
		t.Errorf("Expected size 1000, got %d", q.Size())
	}
}

func TestConcurrentPopFirst(t *testing.T) {
	q := NewDequeConcurrency[int]()

	// Fill the deque first
	for i := 0; i < 1000; i++ {
		q.PushLast(i)
	}

	var wg sync.WaitGroup
	successCount := make(chan int, 1000) // using a channel to count successful pops

	// PopFirst concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := q.PopFirst()
			if err == nil {
				successCount <- 1
			}
		}()
	}

	wg.Wait()
	close(successCount)

	count := 0
	for range successCount {
		count++
	}

	if count != 1000 {
		t.Errorf("Expected 1000 successful pops, got %d", count)
	}
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
}

func TestConcurrentPopLast(t *testing.T) {
	q := NewDequeConcurrency[int]()

	// Fill the deque first
	for i := 0; i < 1000; i++ {
		q.PushFirst(i)
	}

	var wg sync.WaitGroup
	successCount := make(chan int, 1000) // using a channel to count successful pops

	// PopLast concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := q.PopLast()
			if err == nil {
				successCount <- 1
			}
		}()
	}

	wg.Wait()
	close(successCount)

	count := 0
	for range successCount {
		count++
	}

	if count != 1000 {
		t.Errorf("Expected 1000 successful pops, got %d", count)
	}
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
}

func TestGetFirst(t *testing.T) {
	dq := NewDequeConcurrency[int]()
	dq.PushLast(1)
	dq.PushLast(2)
	dq.PushLast(3)

	ch := make(chan int)

	go func() {
		val, err := dq.GetFirst()
		if err != nil {
			t.Errorf("Error while getting first element: %v", err)
		}
		ch <- val
	}()

	select {
	case val := <-ch:
		if val != 1 {
			t.Errorf("Expected first element to be 1, got %v", val)
		}
	}
}

func TestGetLast(t *testing.T) {
	dq := NewDequeConcurrency[int]()
	dq.PushLast(1)
	dq.PushLast(2)
	dq.PushLast(3)

	ch := make(chan int)

	go func() {
		val, err := dq.GetLast()
		if err != nil {
			t.Errorf("Error while getting last element: %v", err)
		}
		ch <- val
	}()

	select {
	case val := <-ch:
		if val != 3 {
			t.Errorf("Expected last element to be 3, got %v", val)
		}
	}
}
