package concurrency

import (
	"sync"

	"github.com/a1ex3/go-containers/containers/queue"
)

type queueConcurrency[T any] struct {
	mu         sync.Mutex
	queueSync_ *queue.Queue[T]
}

type QueueConcurrency[T any] struct {
	queue_ *queueConcurrency[T]
}

func newQueueConcurrency[T any]() *queueConcurrency[T] {
	queue := queue.NewQueue[T]()
	return &queueConcurrency[T]{
		mu:         sync.Mutex{},
		queueSync_: queue,
	}
}

func NewQueueConcurrency[T any]() *QueueConcurrency[T] {
	qe := newQueueConcurrency[T]()
	return &QueueConcurrency[T]{
		queue_: qe,
	}
}
