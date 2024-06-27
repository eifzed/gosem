package main

import (
	"fmt"
	"time"

	"github.com/eifzed/gosem"
)

func main() {
	semaphore := gosem.NewSemaphore(gosem.WithMaxWorker(2), gosem.WithTimeout(5), gosem.WithPanicHandler(panicHandler))
	semaphore.SetFunc(Foo) // setting the function

	// a list of dummy data to be processed
	data := [][]int{{0, 1}, {1, 3}, {2, 10}, {3, 1}}

	for _, d := range data {
		err := semaphore.Call(d[0], d[1])
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	semaphore.Close()

	// do other things
}
func panicHandler() {
	if r := recover(); r != nil {
		fmt.Printf("Recovered from panic: %v\n", r)
	}
}

func Foo(id int, delaySecond int) error {
	fmt.Println("start id", id, "delay", delaySecond, "second")
	time.Sleep(time.Duration(delaySecond) * time.Second)
	fmt.Println("end id", id)
	return nil
}
