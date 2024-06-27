package gosem

func NewSemaphore(opts ...OptFunc) *Semaphore {
	s := defaultOpts()

	for _, fn := range opts {
		fn(s)
	}

	return s
}

func (s *Semaphore) SetPanicHandler(fn func()) {
	s.hasPanicHandler = true
	s.panicHandler = fn
}

func (s *Semaphore) SetTimeout(timeoutSecond uint) {
	s.hasTimeout = true
	s.timeoutSecond = timeoutSecond
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}
