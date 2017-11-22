package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// a blocking channel to keep concurrency under control
	semaphoreChan := make(chan struct{}, 10)
	defer close(semaphoreChan)

	// a wait group enables the main process a wait for goroutines to finish
	wg := sync.WaitGroup{}

	// a simple loop from 1 to 10
	for i := 1; i <= 100; i++ {

		// increment the wait group internal counter
		wg.Add(1)

		// print what we're about to be doing (will keep the order)
		fmt.Printf("About to run #%d in a goroutine\n", i)

		// fire off a goroutine with the index in a closure since it will be modified
		go func(i int) {

			// block until the semaphore channel has room
			// this could also be moved out of the goroutine
			// which would make sense if the list is huge
			semaphoreChan <- struct{}{}

			// pretend to do some synchronous work
			time.Sleep(time.Second)

			// tell the wait group that we be done
			wg.Done()

			// print an message containing the index (won't keep order)
			fmt.Printf("About to exit #%d from a goroutine\n", i)

			// clear a spot in the semaphore channel
			<-semaphoreChan

		}(i)
	}

	// wait for all the goroutines to be done
	wg.Wait()
}
