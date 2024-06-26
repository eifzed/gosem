package gosem

import "reflect"

type OptFunc func(*Worker)

type Worker struct {
	fn               reflect.Value
	semaphoreChannel chan struct{}
	timeoutSecond    uint
	hasTimeout       bool
	panicHandler     func()
	hasPanicHandler  bool
	workerCount      uint
}
