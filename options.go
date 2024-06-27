package gosem

import (
	"fmt"
	"sync"
)

// defaultOpts set default options for the semaphore
// does not use timeout, max worker is set to 2
func defaultOpts() *Semaphore {
	return &Semaphore{
		hasTimeout:       false,
		semaphoreChannel: make(chan struct{}, 2),
		timeoutSecond:    0,
		wg:               &sync.WaitGroup{},
	}
}

// WithTimeout sets the timeout of the worker
func WithTimeout(timeoutSecond uint) OptFunc {
	return func(s *Semaphore) {
		s.timeoutSecond = timeoutSecond
		s.hasTimeout = true
	}
}

// WithMaxWorker setes the maximum number of worker
// that's allowed to be spawed
func WithMaxWorker(maxWorker uint) OptFunc {
	return func(s *Semaphore) {
		if maxWorker < 2 {
			maxWorker = 2
		}
		s.semaphoreChannel = make(chan struct{}, maxWorker)
	}
}

// WithDefaultPanicHandler uses a simple panic handler
func WithDefaultPanicHandler() OptFunc {
	return func(s *Semaphore) {
		s.hasPanicHandler = true
		s.panicHandler = func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}
	}
}

// WithDefaultPanicHandler uses a custom panic handler
func WithPanicHandler(fn func()) OptFunc {
	return func(s *Semaphore) {
		s.hasPanicHandler = true
		s.panicHandler = fn
	}
}
