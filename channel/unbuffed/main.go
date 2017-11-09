package main

import (
	"fmt"
	"time"
)

// https://nanxiao.gitbooks.io/golang-101-hacks/content/posts/unbuffered-and-buffered-channels.html
// After the main goroutine is launched, it will sleep immediately("Main goroutine sleeps 2 seconds" is printed),
// and this will cause main goroutine relinquishes the CPU to the func goroutine("Func goroutine begins sending data" is printed).
// But since the main goroutine is sleeping and can't receive data from the channel,
// so ch <- 1 operation in func goroutine can't complete until d := <- ch in main goroutine is executed(The final 3 logs are printed).
func main() {

	ch := make(chan int)

	go func(ch chan int) {
		fmt.Println("Func goroutine begins sending data")
		// sender will block on the channel until the receiver receives the data from the channel
		ch <- 1
		fmt.Println("Func goroutine ends sending data")
	}(ch)

	fmt.Println("Main goroutine sleeps 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("Main goroutine begins receiving data")
	// the receiver will also block on the channel until sender sends data into the channel.
	d := <-ch
	fmt.Println("Main goroutine received data:", d)

	time.Sleep(time.Second)
}
