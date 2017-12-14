package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	for i := 0; i <= 10; i++ {
		go print(i)
		wg.Add(1)
	}
	wg.Wait()
}

func looprint(count int) {
	for i := 0; i <= count; i++ {
		fmt.Println(i)
	}
}

func print(i int) {
	fmt.Println(i)
	wg.Done()
}
