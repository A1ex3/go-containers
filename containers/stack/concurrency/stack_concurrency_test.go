package concurrency

import (
	"sync"
	"testing"

	"github.com/a1ex3/go-containers/containers/stack"
)

func TestStackConcurrency_PushPopConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	testData := make([]int, 0)
	var stack stack.IStack[int] = NewStackConcurrency[int]()

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
