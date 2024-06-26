package gosem

import "fmt"

func defaultOpts() *Worker {
	return &Worker{
		hasTimeout:       false,
		semaphoreChannel: make(chan struct{}, 2),
		timeoutSecond:    0,
	}
}

func WithTimeout(timeoutSecond uint) OptFunc {

	return func(w *Worker) {
		w.timeoutSecond = timeoutSecond
		w.hasTimeout = true
	}
}

func WithMaxWorker(maxWorker uint) OptFunc {
	return func(w *Worker) {
		if maxWorker < 2 {
			maxWorker = 2
		}
		w.semaphoreChannel = make(chan struct{}, maxWorker)
	}
}

func WithDefaultPanicWrapper() OptFunc {
	return func(w *Worker) {
		w.hasPanicHandler = true
		w.panicHandler = func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}
	}

}

func WithPanicHandler(fn func()) OptFunc {
	return func(w *Worker) {
		w.hasPanicHandler = true
		w.panicHandler = fn
	}
}
