package concurrency

import (
	"sync"
	"testing"

	"github.com/a1ex3/go-containers/containers/stack"
)

func TestStackConcurrency_PushPopConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	testData := make([]int, 0)
	var stack stack.IStack[int] = NewStackDequeConcurrency[int]()

	for i := 0; i < 1000; i++ {
		testData = append(testData, i)
	}

	for _, v := range testData {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			stack.Push(v)
		}(v)
	}
	wg.Wait()

	if stack.Size() != len(testData) {
		t.Errorf("Expected stack length: %d, current stack length: %d", stack.Size(), len(testData))
	}

	stackSize := stack.Size()
	for i := 0; i < stackSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := stack.Pop(); err != nil {
				t.Error(err.Error())
			}
		}()
	}
	wg.Wait()

	if sS := stack.Size(); sS != 0 {
		t.Errorf("Expected stack length: 0, current stack length: %d", sS)
	}
}

func TestTop(t *testing.T) {
	s := NewStackDequeConcurrency[int]()
	var wg sync.WaitGroup

	// Number of goroutines to spawn
	const numGoroutines = 10

	// Synchronize the starting of all goroutines
	start := make(chan struct{})

	// Push elements concurrently
	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			<-start
			s.Push(val)
		}(i)
	}

	// Signal all goroutines to start
	close(start)
	wg.Wait()

	// Test Top on non-empty stack after all pushes
	top, err := s.Top()
	if err != nil {
		t.Errorf("Did not expect an error when calling Top on a non-empty stack: %v", err)
	}
	if top < 1 || top > numGoroutines {
		t.Errorf("Expected top element to be within the range 1 to %d, got %v", numGoroutines, top)
	}

	// Synchronize the starting of all goroutines
	start = make(chan struct{})

	// Pop elements concurrently
	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			s.Pop()
		}()
	}

	// Signal all goroutines to start
	close(start)
	wg.Wait()

	// Test Top on empty stack
	if _, err := s.Top(); err == nil {
		t.Errorf("Expected an error when calling Top on an empty stack")
	}
}

func TestStackSliceConcurrency_PushPopConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	testData := make([]int, 0)
	var stack stack.IStack[int] = NewStackSliceConcurrency[int]()

	for i := 0; i < 1000; i++ {
		testData = append(testData, i)
	}

	for _, v := range testData {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			stack.Push(v)
		}(v)
	}
	wg.Wait()

	if stack.Size() != len(testData) {
		t.Errorf("Expected stack length: %d, current stack length: %d", stack.Size(), len(testData))
	}

	stackSize := stack.Size()
	for i := 0; i < stackSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := stack.Pop(); err != nil {
				t.Error(err.Error())
			}
		}()
	}
	wg.Wait()

	if sS := stack.Size(); sS != 0 {
		t.Errorf("Expected stack length: 0, current stack length: %d", sS)
	}
}

func TestTopSlice(t *testing.T) {
	s := NewStackSliceConcurrency[int]()
	var wg sync.WaitGroup

	// Number of goroutines to spawn
	const numGoroutines = 10

	// Synchronize the starting of all goroutines
	start := make(chan struct{})

	// Push elements concurrently
	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			<-start
			s.Push(val)
		}(i)
	}

	// Signal all goroutines to start
	close(start)
	wg.Wait()

	// Test Top on non-empty stack after all pushes
	top, err := s.Top()
	if err != nil {
		t.Errorf("Did not expect an error when calling Top on a non-empty stack: %v", err)
	}
	if top < 1 || top > numGoroutines {
		t.Errorf("Expected top element to be within the range 1 to %d, got %v", numGoroutines, top)
	}

	// Synchronize the starting of all goroutines
	start = make(chan struct{})

	// Pop elements concurrently
	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			s.Pop()
		}()
	}

	// Signal all goroutines to start
	close(start)
	wg.Wait()

	// Test Top on empty stack
	if _, err := s.Top(); err == nil {
		t.Errorf("Expected an error when calling Top on an empty stack")
	}
}
