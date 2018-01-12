// This sample program demonstrates how to use the atomic
// package to provide safe access to numeric types.

// Atomic functions provide low-level locking mechanisms for synchronizing access to
// integers and pointers.

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64

	wg sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()
	fmt.Println("Final Counter : ", counter)
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
