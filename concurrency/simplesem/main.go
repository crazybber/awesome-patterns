package main

import (
	"fmt"
	"time"
)

func main() {
	concurrentCount := 5
	sem := make(chan bool, concurrentCount)

	for i := 0; i < 100; i++ {
		sem <- true
		go func(i int) {
			dummyProcess(i)
			<-sem
		}(i)
	}
	time.Sleep(time.Second * 20)
}

func dummyProcess(i int) {
	fmt.Printf("Process Dummy %v \n", i)
	time.Sleep(time.Second)
}
