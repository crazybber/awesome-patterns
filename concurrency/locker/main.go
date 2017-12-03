package main

/**
In a particularly sensitive portion of code, you need to lock certain resources. Given
the frequent use of channels in your code, you’d like to do this with channels instead
of the sync package.

When talking about using a channel as a lock, you want this kind of behavior:
1 A function acquires a lock by sending a message on a channel.
2 The function proceeds to do its sensitive operations.
3 The function releases the lock by reading the message back off the channel.
4 Any function that tries to acquire the lock before it’s been released will pause
when it tries to acquire the (already locked) lock.
**/
import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool, 2)
	for i := 1; i < 7; i++ {
		go worker(i, lock)
	}
	time.Sleep(time.Second * 10)
}

func worker(id int, lock chan bool) {
	fmt.Printf("%d wants the lock\n", id)
	lock <- true
	fmt.Printf("%d has the lock\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d is releasing the lock\n", id)
	<-lock
}
