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
