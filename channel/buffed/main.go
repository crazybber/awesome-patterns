package main

import (
	"fmt"
	"time"
)

// Compared with unbuffered counterpart, the sender of buffered channel will block when there is no empty slot of the channel,
// while the receiver will block on the channel when it is empty.
func main() {
	ch := make(chan int, 2)

	go func(ch chan int) {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Println("Func goroutine sends data: ", i)
		}
		close(ch)
	}(ch)

	fmt.Println("Main goroutine sleeps 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("Main goroutine begins receiving data")
	for d := range ch {
		fmt.Println("Main goroutine received data:", d)
	}
}

// Main goroutine sleeps 2 seconds

// Func goroutine sends data: 1
// Main goroutine received data: 1

// Func goroutine sends data: 2
// Main goroutine received data: 2

// Main goroutine begins receiving data
// Func goroutine sends data: 3
// Main goroutine received data: 3
