package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestMergeNumbersSeq(T *testing.T) {
	randomNumbers := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	// generate the common channel with inputs
	inputChan := generatePipeline(randomNumbers)

	// Fan-out to 2 Go-routine
	c1 := squareNumber(inputChan)
	c2 := squareNumber(inputChan)

	// Fan-in the resulting squared numbers
	c := fanIn(c1, c2)
	sum := 0

	// Do the summation
	for i := 0; i < len(randomNumbers); i++ {
		sum += <-c
	}
	fmt.Printf("Total Sum of Squares: %d", sum)
}

func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func squareNumber(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// fanIn compose different channels into one
func fanIn(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int, 3)

	wg.Add(len(cs))

	// Start an send goroutine for each input channel in cs. send
	// copies values from c to out until c is closed, then calls wg.Done.
	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	//启动多个 go routine 开始工作
	for _, c := range cs {
		go send(c)
	}
	// Start a goroutine to close out once all the send goroutines are
	// done.  This must start after the wg.Add call.
	//关闭动作,放在发送一方，会更好
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
