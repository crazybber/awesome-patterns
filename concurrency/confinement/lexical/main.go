package main

import "fmt"

// Confinement is the simple yet powerful idea of ensuring information is only ever available from one concurrent process.
// There are two kinds of confinement possible: ad hoc and lexical.

// Lexical confinement involves using lexical scope to expose only the correct data and
// concurrency primitives for multiple concurrent processes to use. It makes it impossible to do the wrong thing.

func main() {
	lexicalDemo()
}

func lexicalDemo() {
	// Here we instantiate the channel within the lexical scope of the chanOwner function.
	// This limits the scope of the write aspect of the results channel to the closure
	// defined below it. In other words, it confines the write aspect of this channel to
	// prevent other goroutines from writing to it.
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// Here we receive a read-only copy of an int channel. By declaring that the only
	// usage we require is read access, we confine usage of the channel within the consume function to only reads
	comsumer := func(results <-chan int) {
		for result := range results {
			fmt.Println("Received: %d\n", result)
		}
		fmt.Println("Done Receiving!")
	}

	// Here we receive the read aspect of the channel and we’re able to pass it into the
	// consumer, which can do nothing but read from it. Once again this confines the
	// main goroutine to a read-only view of the channel.
	results := chanOwner()
	comsumer(results)
}
