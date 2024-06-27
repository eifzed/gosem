package gosem

import (
	"reflect"
	"sync"
)

type OptFunc func(*Semaphore)

type Semaphore struct {
	fn               reflect.Value
	semaphoreChannel chan struct{}
	timeoutSecond    uint
	hasTimeout       bool
	panicHandler     func()
	hasPanicHandler  bool
	workerCount      uint
	wg               *sync.WaitGroup
}
