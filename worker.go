package gosem

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"
)

var (
	ErrInvalidFunction    = errors.New("parameter should be a type of function")
	ErrUnsetFunction      = errors.New("Set() function must be used prior to using Call() function")
	ErrMismatchParameters = errors.New("the number of parameters does not match the function signature")
	// ErrTimeout            = errors.New("the number of parameters does not match the function signature")
)

// SetFunc sets the function to be executed
func (w *Semaphore) SetFunc(function interface{}) error {
	fn := reflect.ValueOf(function)
	if fn.Kind() != reflect.Func {
		return ErrInvalidFunction
	}
	w.fn = fn
	return nil
}

// SetFunc calls the function.
// Must use SetFunc() first before using this function
func (w *Semaphore) Call(params ...interface{}) error {
	if !w.fn.IsValid() {
		return ErrUnsetFunction
	}
	if len(params) != w.fn.Type().NumIn() {
		return ErrMismatchParameters
	}
	args := make([]reflect.Value, len(params))
	for i, param := range params {
		args[i] = reflect.ValueOf(param)
	}

	w.semaphoreChannel <- struct{}{}

	go w.execute(args...)

	return nil
}

func (w *Semaphore) execute(args ...reflect.Value) {
	defer func() {
		<-w.semaphoreChannel
	}()

	var ctx context.Context

	if w.hasTimeout {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(w.timeoutSecond)*time.Second)
		defer cancel()
	} else {
		ctx = context.Background()
	}
	w.doCall(ctx, args...)

}

func (w *Semaphore) doCall(ctx context.Context, args ...reflect.Value) {
	if w.hasPanicHandler {
		defer w.panicHandler()
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Function execution stopped due to timeout")
			return
		default:
			w.fn.Call(args)
		}
	}
}
