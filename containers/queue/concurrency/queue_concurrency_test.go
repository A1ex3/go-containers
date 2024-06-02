package concurrency

import (
	"sync"
	"testing"
)

func TestQueue_PushConcurrently(t *testing.T) {
	q := NewQueueDequeConcurrency[int]()
	var wg sync.WaitGroup
	const numPushers = 2
	wg.Add(numPushers)
	for i := 0; i < numPushers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				q.Push(j)
			}
		}()
	}
	wg.Wait()
	if q.Size() != 2000 {
		t.Errorf("Expected queue size to be 2000, got %d", q.Size())
	}
}

func TestQueue_PopConcurrently(t *testing.T) {
	q := NewQueueDequeConcurrency[int]()
	for i := 0; i < 1000; i++ {
		q.Push(i)
	}
	var wg sync.WaitGroup
	const numPoppers = 2
	wg.Add(numPoppers)
	for i := 0; i < numPoppers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 500; j++ {
				q.Pop()
			}
		}()
	}
	wg.Wait()
	if q.Size() != 0 {
		t.Errorf("Expected queue size to be 0, got %d", q.Size())
	}
}

func TestFront(t *testing.T) {
	q := NewQueueDequeConcurrency[int]() // Create a new instance of the queue.

	// Add elements to the queue
	for i := 0; i < 100; i++ {
		q.Push(i)
	}

	var wg sync.WaitGroup
	wg.Add(100)

	// Concurrently call Front() method
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			front, err := q.Front()
			if err != nil {
				t.Errorf("Error getting front element: %v", err)
			}
			// Do something with front element if needed
			_ = front
		}()
	}

	wg.Wait() // Wait for all goroutines to finish
}

func TestQueueSlice_PushConcurrently(t *testing.T) {
	q := NewQueueSliceConcurrency[int]()
	var wg sync.WaitGroup
	const numPushers = 2
	wg.Add(numPushers)
	for i := 0; i < numPushers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				q.Push(j)
			}
		}()
	}
	wg.Wait()
	if q.Size() != 2000 {
		t.Errorf("Expected queue size to be 2000, got %d", q.Size())
	}
}

func TestQueueSlice_PopConcurrently(t *testing.T) {
	q := NewQueueSliceConcurrency[int]()
	for i := 0; i < 1000; i++ {
		q.Push(i)
	}
	var wg sync.WaitGroup
	const numPoppers = 2
	wg.Add(numPoppers)
	for i := 0; i < numPoppers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 500; j++ {
				q.Pop()
			}
		}()
	}
	wg.Wait()
	if q.Size() != 0 {
		t.Errorf("Expected queue size to be 0, got %d", q.Size())
	}
}

func TestFrontSlice(t *testing.T) {
	q := NewQueueSliceConcurrency[int]() // Create a new instance of the queue.

	// Add elements to the queue
	for i := 0; i < 100; i++ {
		q.Push(i)
	}

	var wg sync.WaitGroup
	wg.Add(100)

	// Concurrently call Front() method
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			front, err := q.Front()
			if err != nil {
				t.Errorf("Error getting front element: %v", err)
			}
			// Do something with front element if needed
			_ = front
		}()
	}

	wg.Wait() // Wait for all goroutines to finish
}
