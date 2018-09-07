package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestFiveGopherWithChan(t *testing.T) {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopherWithChan(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}

func sleepyGopherWithChan(id int, c chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("... ", id, " snore ...")
	c <- id
}
