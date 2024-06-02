package concurrency

import (
	"sync"

	"github.com/a1ex3/go-containers/containers/queue"
)

type QueueDequeConcurrency[T any] struct {
	mu         sync.Mutex
	queueSync_ *queue.QueueDeque[T]
}

type QueueSliceConcurrency[T any] struct {
	mu         sync.Mutex
	queueSync_ *queue.QueueSlice[T]
}

func NewQueueDequeConcurrency[T any]() *QueueDequeConcurrency[T] {
	return &QueueDequeConcurrency[T]{
		mu:         sync.Mutex{},
		queueSync_: queue.NewQueueDeque[T](),
	}
}

func NewQueueSliceConcurrency[T any]() *QueueSliceConcurrency[T] {
	return &QueueSliceConcurrency[T]{
		mu:         sync.Mutex{},
		queueSync_: queue.NewQueueSlice[T](),
	}
}