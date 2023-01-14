package concurrency

func (p *Pool[T]) Enqueue(job func(params ...T), params ...T) {
	if p.jobs == 0 {
		go job(params...)
		return
	}
	p.queue <- struct{}{}
	go func() {
		defer func() {
			<-p.queue
		}()
		job(params...)
	}()
}
