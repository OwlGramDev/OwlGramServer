package concurrency

func NewPool[T any](size int) *Pool[T] {
	if size < 0 {
		size = 0
	}
	return &Pool[T]{
		jobs:  size,
		queue: make(chan struct{}, size),
	}
}
