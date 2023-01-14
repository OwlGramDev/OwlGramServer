package concurrency

type Pool[T any] struct {
	jobs  int
	queue chan struct{}
}
