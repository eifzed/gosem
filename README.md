# gosem

`gosem` is a Go concurrency library helper that provides a simple and efficient way to control the number of concurrently running goroutines using the semaphore pattern. This library is most useful for when you have a list of data that needs to be executed individually and concurrently with one function.

## Features

- Easy-to-use semaphore pattern implementation.
- Control the number of concurrently running goroutines.
- Control the timeout for each worker.
- Set a default or custom Panic handler

## Installation

To install `gosem`, use `go get`:

```sh
go get github.com/eifzed/gosem
```

## Usage

Here's a basic example demonstrating how to use `gosem`:

```go
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

	// waits all workers to finish
	semaphore.Wait()

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
	if id == 3 {
		// dummy panic
		panic("abc")
	}
	return nil
}




```

### API

#### `NewSemaphore`

Creates a new semaphore with options to set the max number of workers, timeout, etc. if options are not provided then default values are used.

```go
func NewSemaphore(WithTimeout(5), WithMaxWorker(2)) *Semaphore
```

- `WithTimeout`: Sets the maximum amount of time the function is allowed to execute.
- `WithMaxWorker`: Sets the maximum number of worker to execute the functions.

#### `SetFunc`

Sets the function to be executed
```go
func (s *Semaphore) SetFunc(fn func())
```

#### `Call`

Calls the function
```go
func (s *Semaphore) Call(args ...interface{})
```

#### `Wait`

Waits all workers to finish executing the function
```go
func (s *Semaphore) Wait()
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have any improvements or suggestions.

## Contact

For any questions or issues, please open an issue on this repository or contact the maintainer.

---

Happy concurrency with `gosem`!
