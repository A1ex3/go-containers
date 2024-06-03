package concurrency

import (
	"sync"

	linkedlist "github.com/a1ex3/go-containers/containers/linked_list"
)

type linkedListIteratorConcurrency[T any] struct {
	mu       sync.Mutex
	iterSync *linkedlist.LinkedListIterator[T]
}

func (linkedList *LinkedList[T]) NewIteratorConcurrency() *linkedListIteratorConcurrency[T] {
	iter := linkedList.linkedList_.linkedListSync.NewIterator()
	return &linkedListIteratorConcurrency[T]{
		mu:       sync.Mutex{},
		iterSync: iter,
	}
}

func (it *linkedListIteratorConcurrency[T]) HasNext() bool {
	it.mu.Lock()
	defer it.mu.Unlock()

	return it.iterSync.HasNext()
}

func (it *linkedListIteratorConcurrency[T]) Next() (T, error) {
	it.mu.Lock()
	defer it.mu.Unlock()

	return it.iterSync.Next()
}
