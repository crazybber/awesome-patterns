package main

//Multiple functions can read from the same channel until that channel is closed;
// this is called fan-out. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.

// A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels
// onto a single channel that's closed when all the inputs are closed. This is called fan-in.

import (
	"fmt"
	"sync"
)

/*
There's no formal definition of a pipeline in Go; it's just one of many kinds of concurrent programs. Informally, a pipeline is a series of stages connected by channels, where each stage is a group of goroutines running the same function. In each stage, the goroutines

receive values from upstream via inbound channels
perform some function on that data, usually producing new values
send values downstream via outbound channels
Each stage has any number of inbound and outbound channels, except the first and last stages, which have only outbound or inbound channels, respectively. The first stage is sometimes called the source or producer; the last stage, the sink or consumer.
*/
func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

//The merge function converts a list of channels to a single channel by starting a goroutine for each inbound channel that copies
// the values to the sole outbound channel. Once all the output goroutines have been started,
// merge starts one more goroutine to close the outbound channel after all sends on that channel are done.

// Sends on a closed channel panic, so it's important to ensure all sends are done before calling close.
// The sync.WaitGroup type provides a simple way to arrange this synchronization:
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

//The first stage, gen, is a function that converts a list of integers to a channel that emits the integers in the list.
// The gen function starts a goroutine that sends the integers on the channel and closes the channel when all
// the values have been sent:
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

//The second stage, sq, receives integers from a channel and returns a channel that emits the square of
// each received integer. After the inbound channel is closed and this stage has sent all the values downstream,
// it closes the outbound channel:
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
