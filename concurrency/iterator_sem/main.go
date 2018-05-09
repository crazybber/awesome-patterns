package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	generator := func() <-chan int {
		totalIteration := 30
		wg := &sync.WaitGroup{}
		wg.Add(totalIteration)
		concurrentCount := 20
		sem := make(chan bool, concurrentCount)
		defer close(sem)
		results := make(chan int, totalIteration)
		defer close(results)
		for i := 0; i < totalIteration; i++ {
			sem <- true
			go func(i int) {
				results <- generateNumber()
				<-sem
				wg.Done()
			}(i)
		}
		wg.Wait()
		return results
	}

	results := generator()
	for r := range results {
		fmt.Println(r)
	}

	fmt.Println("Finished!!!")
}

func generateNumber() int {
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Second)
	return rand.Intn(10000)
}
