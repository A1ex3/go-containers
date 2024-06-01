package stack

import (
	"sort"
	"testing"
)

func TestStack_PushPop(t *testing.T) {
	var testData []int = []int{1, 2, 3, 4, 5}
	var stack IStack[int] = NewStack[int]()

	for _, v := range testData {
		stack.Push(v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(testData)))
	for i := 0; i < stack.Size(); i++ {
		testDataVal := testData[i]

		if val, err := stack.Pop(); val != testDataVal || err != nil {
			t.Errorf("Test data: %d, Stack data: %d. Failed.", testDataVal, val)
		}
	}
}

func TestTop(t *testing.T) {
	s := NewStack[int]()

	// Test Top on empty stack
	if _, err := s.Top(); err == nil {
		t.Errorf("Expected an error when calling Top on an empty stack")
	}

	// Push elements to the stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test Top on non-empty stack
	if top, err := s.Top(); err != nil {
		t.Errorf("Did not expect an error when calling Top on a non-empty stack: %v", err)
	} else if top != 3 {
		t.Errorf("Expected top element to be 3, got %v", top)
	}

	// Pop the top element and check again
	s.Pop()
	if top, err := s.Top(); err != nil {
		t.Errorf("Did not expect an error when calling Top on a non-empty stack: %v", err)
	} else if top != 2 {
		t.Errorf("Expected top element to be 2, got %v", top)
	}

	// Pop all elements and test Top on empty stack again
	s.Pop()
	s.Pop()
	if _, err := s.Top(); err == nil {
		t.Errorf("Expected an error when calling Top on an empty stack")
	}
}
