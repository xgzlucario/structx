package structx

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type Pool[T any] struct {
	sem  chan struct{} // limit goroutine
	work chan poolTask[T]
	len  *int32
	wg   sync.WaitGroup
}

type poolTask[T any] struct {
	work   func(...T)
	params []T
}

// NewPool: Return new pool
func NewPool[T any](size ...int) *Pool[T] {
	// default
	num := runtime.NumCPU()
	if len(size) > 0 {
		num = size[0]
	}
	return &Pool[T]{
		work: make(chan poolTask[T]),
		sem:  make(chan struct{}, num),
		len:  new(int32),
	}
}

// NewTask: Submit New Task
func (p *Pool[T]) NewTask(task func(...T), params ...T) {
	p.wg.Add(1)
	t := poolTask[T]{
		work:   task,
		params: params,
	}
	select {
	case p.work <- t:
		atomic.AddInt32(p.len, 1)

	case p.sem <- struct{}{}:
		go p.worker(t)
		atomic.AddInt32(p.len, 1)
	}
}

// Do Task Backend
func (p *Pool[T]) worker(t poolTask[T]) {
	defer func() { <-p.sem }()
	var ok = true
	for ok {
		t.work(t.params...)
		p.wg.Done()
		atomic.AddInt32(p.len, -1)
		t, ok = <-p.work
	}
}

// Wait
func (p *Pool[T]) Wait() {
	p.wg.Wait()
}

// Len
func (p *Pool[T]) Len() int32 {
	return atomic.LoadInt32(p.len)
}

// Close
func (p *Pool[T]) Close() {
	close(p.work)
	close(p.sem)
}
