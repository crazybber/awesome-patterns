package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestOneGopher(t *testing.T) {
	go sleepyGopher(1)
	time.Sleep(3 * time.Second)
}

func TestFiveGopher(t *testing.T) {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopherWithChan(i, c)
	}
	time.Sleep(3 * time.Second)
}

func sleepyGopher(id int) {
	time.Sleep(2 * time.Second)
	fmt.Println("... ", id, " snore ...")
}
