package main

import "github.com/davecgh/go-spew/spew"

func main() {
	run()
}

func run() {
	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4)
	pipline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipline {
		spew.Dump(v)
	}
}

func generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}

func multiply(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case multipliedStream <- i * multiplier:
			}
		}
	}()
	return multipliedStream
}

func add(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
	addStream := make(chan int)
	go func() {
		defer close(addStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case addStream <- i:
			}
		}
	}()
	return addStream
}
