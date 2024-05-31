package iterator

type IIterator[T any] interface {
	HasNext() bool
	Next() (T, error)
}
