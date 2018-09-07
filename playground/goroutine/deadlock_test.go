package goroutine

import "testing"

// A deadlock happens when a group of goroutines are waiting for each other and none of them is able to proceed.
// The program will get stuck on the channel send operation waiting forever for someone to read the value.
// Go is able to detect situations like this at runtime.
func TestDeadlock(t *testing.T) {
	c := make(chan int)
	<-c
}
