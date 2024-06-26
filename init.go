package gosem

func NewWorker(opts ...OptFunc) *Worker {
	w := defaultOpts()

	for _, fn := range opts {
		fn(w)
	}

	return w
}

func (w *Worker) SetPanicHandler(fn func()) {
	w.hasPanicHandler = true
	w.panicHandler = fn
}

func (w *Worker) SetTimeout(timeoutSecond uint) {
	w.hasTimeout = true
	w.timeoutSecond = timeoutSecond
}
