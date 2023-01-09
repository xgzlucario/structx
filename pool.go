package structx

import (
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/sourcegraph/conc"
)

// Pool
type Pool struct {
	handle   conc.WaitGroup
	limiter  limiter
	tasks    chan func()
	initOnce sync.Once
	len      *int32
}

// NewPool
func NewPool() *Pool {
	return &Pool{}
}

// Go submits a task to be run in the pool.
func (p *Pool) Go(f func()) {
	p.init()

	select {
	case p.limiter <- struct{}{}:
		p.handle.Go(p.worker)
		p.tasks <- f
		atomic.AddInt32(p.len, 1)

	case p.tasks <- f:
		atomic.AddInt32(p.len, 1)
	}
}

// Wait
func (p *Pool) Wait() {
	p.init()

	close(p.tasks)
	p.handle.Wait()
}

// MaxGoroutines
func (p *Pool) MaxGoroutines() int {
	return cap(p.limiter)
}

// Len: Running tasks num
func (p *Pool) Len() int32 {
	return atomic.LoadInt32(p.len)
}

// WithMaxGoroutines
func (p *Pool) WithMaxGoroutines(n int) *Pool {
	if n < 1 {
		panic("max goroutines in a pool must be greater than zero")
	}
	p.limiter = make(limiter, n)
	return p
}

// init
func (p *Pool) init() {
	p.initOnce.Do(func() {
		// Do not override the limiter if set by WithMaxGoroutines
		if p.limiter == nil {
			p.limiter = make(limiter, runtime.GOMAXPROCS(0))
		}

		p.len = new(int32)
		p.tasks = make(chan func())
	})
}

func (p *Pool) worker() {
	defer func() { <-p.limiter }()

	for f := range p.tasks {
		f()
		atomic.AddInt32(p.len, -1)
	}
}

type limiter chan struct{}
