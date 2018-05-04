package main

import "fmt"

// Confinement is the simple yet powerful idea of ensuring information is only ever
// available from one concurrent process.

// When this is achieved, a concurrent program
//is implicitly safe and no synchronization is needed.

// Ad hoc confinement is when you achieve confinement through a convention—
// whether it be set by the languages community, the group you work within, or the
// codebase you work within
func main() {
	data := make([]int, 4)
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}
	handleData := make(chan int)
	go loopData(handleData)
	for num := range handleData {
		fmt.Println(num)
	}

	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				// Do something interesting
				fmt.Println(s, "TTT")
			}
		}()
		return completed
	}
	doWork(nil)
	// Perhaps more work is done here
	fmt.Println("Done.")
}
