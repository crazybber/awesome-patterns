package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/golang/go/src/pkg/math/rand"
)

// Confinement is the simple yet powerful idea of ensuring information is only ever available from one concurrent process.
// There are two kinds of confinement possible: ad hoc and lexical.

// Lexical confinement involves using lexical scope to expose only the correct data and
// concurrency primitives for multiple concurrent processes to use. It makes it impossible to do the wrong thing.

// If a goroutine is responsible for creating a goroutine, it is also responsible for ensuring it can stop the goroutine.

func main() {
	// lexicalNotConcurrentSafe()
	// lexicalDemo()
	// blockOnAttemptingToWriteToChannel()
	fixBlockOnAttemptingToWriteToChannel()
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

func lexicalNotConcurrentSafe() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	// Here we pass in a slice containing the first three bytes in the data structure.
	go printData(&wg, data[:3])
	// Here we pass in a slice containing the last three bytes in the data structure.
	go printData(&wg, data[3:])
	wg.Wait()
}

func blockOnAttemptingToWriteToChannel() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure existed.")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}
	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}

// The solution, just like for the receiving case, is to provide the
// producer goroutine with a channel informing it to exit
func fixBlockOnAttemptingToWriteToChannel() {
	d := make(chan interface{})
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure existed.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}
	randStream := newRandStream(d)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(d)
	time.Sleep(1 * time.Second)
}
