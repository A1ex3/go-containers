package unorderedset

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	var set IUnorderedSet[int] = NewUnorderedSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	if !set.Contains(1) || !set.Contains(2) || !set.Contains(3) {
		t.Errorf("Add method did not add elements properly %x", set.ToSlice())
	}
}

func TestRemove(t *testing.T) {
	set := NewUnorderedSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Remove(2)

	if set.Contains(2) {
		t.Errorf("Remove method did not remove element properly")
	}
}

func TestContains(t *testing.T) {
	set := NewUnorderedSet[int]()
	set.Add(1)
	set.Add(2)

	if !set.Contains(1) || !set.Contains(2) {
		t.Errorf("Contains method did not return true for existing elements")
	}

	if set.Contains(3) {
		t.Errorf("Contains method returned true for non-existing element")
	}
}

func TestSize(t *testing.T) {
	set := NewUnorderedSet[int]()
	set.Add(1)
	set.Add(2)

	size := set.Size()
	if size != 2 {
		t.Errorf("Size method did not return correct size, expected %d but got %d", 2, size)
	}
}

func TestClear(t *testing.T) {
	set := NewUnorderedSet[int]()
	set.Add(1)
	set.Add(2)
	set.Clear()

	if set.Size() != 0 {
		t.Errorf("Clear method did not clear the set properly")
	}
}

func contains(slice []int, item int) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

func TestToSlice(t *testing.T) {
	set := NewUnorderedSet[int]()
	set.Add(1)
	set.Add(2)

	slice := set.ToSlice()

	if len(slice) != 2 || !contains(slice, 1) || !contains(slice, 2) {
		t.Errorf("ToSlice method did not return correct slice")
	}
}

func TestDifference(t *testing.T) {
	set1 := NewUnorderedSet[int]()
	set2 := NewUnorderedSet[int]()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	set2.Add(2)
	set2.Add(3)
	set2.Add(4)

	diff := set1.Difference(set2)
	expectedDiff := NewUnorderedSet[int]()
	expectedDiff.Add(1)

	if !reflect.DeepEqual(diff, expectedDiff) {
		t.Errorf("Difference failed. Expected %v but got %v", expectedDiff, diff)
	}
}

func TestIntersection(t *testing.T) {
	set1 := NewUnorderedSet[int]()
	set2 := NewUnorderedSet[int]()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	set2.Add(2)
	set2.Add(3)
	set2.Add(4)

	intersection := set1.Intersection(set2)
	expectedIntersection := NewUnorderedSet[int]()
	expectedIntersection.Add(2)
	expectedIntersection.Add(3)

	if !reflect.DeepEqual(intersection, expectedIntersection) {
		t.Errorf("Intersection failed. Expected %v but got %v", expectedIntersection, intersection)
	}
}

func TestSymmetricDifference(t *testing.T) {
	set1 := NewUnorderedSet[int]()
	set2 := NewUnorderedSet[int]()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	set2.Add(2)
	set2.Add(3)
	set2.Add(4)

	symmetricDiff := set1.SymmetricDifference(set2)
	expectedSymmetricDiff := NewUnorderedSet[int]()
	expectedSymmetricDiff.Add(1)
	expectedSymmetricDiff.Add(4)

	if !reflect.DeepEqual(symmetricDiff, expectedSymmetricDiff) {
		t.Errorf("SymmetricDifference failed. Expected %v but got %v", expectedSymmetricDiff, symmetricDiff)
	}
}

func TestUnion(t *testing.T) {
	set1 := NewUnorderedSet[int]()
	set2 := NewUnorderedSet[int]()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)

	set2.Add(2)
	set2.Add(3)
	set2.Add(4)

	union := set1.Union(set2)
	expectedUnion := NewUnorderedSet[int]()
	expectedUnion.Add(1)
	expectedUnion.Add(2)
	expectedUnion.Add(3)
	expectedUnion.Add(4)

	if !reflect.DeepEqual(union, expectedUnion) {
		t.Errorf("Union failed. Expected %v but got %v", expectedUnion, union)
	}
}

func TestIterator(t *testing.T) {
	set := NewUnorderedSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	iterator := set.NewIterator()

	// Test HasNext and Next
	expectedValues := map[int]bool{1: true, 2: true, 3: true}
	remaining := len(expectedValues)

	for iterator.HasNext() {
		nextValue, err := iterator.Next()
		if err != nil {
			t.Errorf("Iterator Next failed: %v", err)
		}

		if _, exists := expectedValues[nextValue]; !exists {
			t.Errorf("Iterator returned unexpected value: %d", nextValue)
		} else {
			delete(expectedValues, nextValue)
			remaining--
		}
	}

	if remaining != 0 {
		t.Errorf("Iterator did not return all expected values. Remaining: %d", remaining)
	}

	// Test Next when there are no more elements
	_, err := iterator.Next()
	if err == nil || err.Error() != "no more elements" {
		t.Errorf("Iterator Next failed. Expected 'no more elements' error, got %v", err)
	}
}
