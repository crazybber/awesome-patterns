package main

import (
	"fmt"
	"time"
)

// demostration of channel and timeout
func main() {
	ourChan := make(chan string, 1)

	go func() {

	}()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Enough waiting")
		close(ourChan)
	}
}
