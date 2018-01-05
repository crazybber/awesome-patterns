// Basic sends and receives on channels are blocking.
// However, we can use `select` with a `default` clause to
// implement _non-blocking_ sends, receives, and even
// non-blocking multi-way `select`s.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(ch chan int) { <-ch }(ch1)
	go func(ch chan int) { ch <- 2 }(ch2)

	time.Sleep(time.Second)

	for {
		select {
		case ch1 <- 1:
			fmt.Println("Send operation on ch1 works!")
		case test := <-ch2:
			fmt.Println("Receive operation on ch2 works!", test)
		default:
			fmt.Println("Exit now!")
			return
		}
	}
}
