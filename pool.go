package structx

import (
	"runtime"
	"sync"
)

type Pool[T any] struct {
	work chan poolTask[T]
	sem  chan struct{} // limit goroutine
	wg   sync.WaitGroup
}

type poolTask[T any] struct {
	work   func(...T)
	params []T
}

// NewPool: Return new pool
func NewPool[T any](size ...int) *Pool[T] {
	// default
	num := runtime.NumCPU() / 2
	if len(size) > 0 {
		num = size[0]
	}
	return &Pool[T]{
		work: make(chan poolTask[T]),
		sem:  make(chan struct{}, num),
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
	case p.sem <- struct{}{}:
		go p.worker(t)
	}
}

// Do Task Forever
func (p *Pool[T]) worker(t poolTask[T]) {
	defer func() { <-p.sem }()
	ok := true
	for ok {
		t.work(t.params...)
		p.wg.Done()
		t, ok = <-p.work
	}
}

func (p *Pool[T]) Wait() {
	defer close(p.work)
	defer close(p.sem)
	p.wg.Wait()
}
