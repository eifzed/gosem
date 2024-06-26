package gosem

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"
)

func (w *Worker) SetFunc(function interface{}) error {
	fn := reflect.ValueOf(function)
	if fn.Kind() != reflect.Func {
		return errors.New("paramter should be a function")
	}
	w.fn = fn
	return nil
}

func (w *Worker) Call(params ...interface{}) error {
	if !w.fn.IsValid() {
		return errors.New("function must be set prior to using Call() by using the Set() function")
	}
	if len(params) != w.fn.Type().NumIn() {
		return errors.New("The number of parameters does not match the function signature")
	}
	args := make([]reflect.Value, len(params))
	for i, param := range params {
		args[i] = reflect.ValueOf(param)
	}

	w.semaphoreChannel <- struct{}{}

	go w.execute(args...)

	return nil
}

func (w *Worker) execute(args ...reflect.Value) {

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

func (w *Worker) doCall(ctx context.Context, args ...reflect.Value) {
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
