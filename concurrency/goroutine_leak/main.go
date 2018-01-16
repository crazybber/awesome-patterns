package main

import "fmt"

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

func cancellationSignal() {

}
