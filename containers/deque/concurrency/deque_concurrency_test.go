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

func TestConcurrentMixedOperations(t *testing.T) {
	q := NewDequeConcurrency[int]()
	var wg sync.WaitGroup
	pushCount := 1000
	popCount := 1000

	// Mixed operations concurrently
	for i := 0; i < pushCount; i++ {
		wg.Add(2) // Increment the wait group counter by 2 for each goroutine
		go func(n int) {
			defer wg.Done()
			q.PushFirst(n)
		}(i)
		go func(n int) {
			defer wg.Done()
			q.PushLast(n)
		}(i)
	}

	successCount := make(chan int, popCount*2)

	for i := 0; i < popCount; i++ {
		wg.Add(2) // Increment the wait group counter by 2 for each goroutine
		go func() {
			defer wg.Done()
			if _, err := q.PopFirst(); err == nil {
				successCount <- 1
			}
		}()
		go func() {
			defer wg.Done()
			if _, err := q.PopLast(); err == nil {
				successCount <- 1
			}
		}()
	}

	// Start a goroutine to wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(successCount)
	}()

	popFirstCount := 0
	for range successCount {
		popFirstCount++
	}

	expectedSize := (pushCount * 2) - popFirstCount
	if popFirstCount != popCount*2 {
		t.Errorf("Expected %d successful pops, got %d", popCount*2, popFirstCount)
	}

	if q.Size() != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, q.Size())
	}
}
