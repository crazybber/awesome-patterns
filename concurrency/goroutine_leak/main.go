package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// The goroutine has a few paths to termination:
// • When it has completed its work.
// • When it cannot continue its work due to an unrecoverable error.
// • When it’s told to stop working

/**
We get the first two paths for free—these paths are your algorithm—but what about
work cancellation? This turns out to be the most important bit because of the net‐
work effect: if you’ve begun a goroutine, it’s most likely cooperating with several other
goroutines in some sort of organized fashion.
**/

func main() {
	cancellationSignal()
}

// Here we see that the main goroutine passes a nil channel into doWork. Therefore, the
// strings channel will never actually gets any strings written onto it, and the goroutine
// containing doWork will remain in memory for the lifetime of this process (we would
// even deadlock if we joined the goroutine within doWork and the main goroutine).
// In this example, the lifetime of the process is very short, but in a real program, gorou‐
// tines could easily be started at the beginning of a long-lived program. In the worst
// case, the main goroutine could continue to spin up goroutines throughout its life,
// causing creep in memory utilization.
func resourceLeak() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// Do something interesting
				fmt.Println(s)
			}
		}()
		return completed
	}
	doWork(nil)
	// Perhaps more work is done here
	fmt.Println("Done.")
}

// The way to successfully mitigate this is to establish a signal between the parent gorou‐
// tine and its children that allows the parent to signal cancellation to its children. By
// convention, this signal is usually a read-only channel named done. The parent gorou‐
// tine passes this channel to the child goroutine and then closes the channel when it
// wants to cancel the child goroutine. Here’s an example:
func cancellationSignal() {
	// Here we pass the done channel to the doWork function. As a convention, this channel is the first parameter.
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				// On this line we see the ubiquitous for-select pattern in use. One of our case statements
				// is checking whether our done channel has been signaled. If it has, we return from the goroutine.
				case t := <-done:
					spew.Dump(t)
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)
	// Here we create another goroutine that will cancel the goroutine spawned in
	// doWork if more than one second passes.
	go func() {
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()
	// This is where we join the goroutine spawned from doWork with the main goroutine.
	<-terminated
	fmt.Println("Done.")
}
