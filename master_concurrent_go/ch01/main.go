package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
}

func main() {
	wg := new(sync.WaitGroup)
	hello := new(Job)
	world := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 3
	world.text = "world"
	world.i = 0
	world.max = 5
	go outputText(hello, wg)
	go outputText(world, wg)
	wg.Add(2)
	wg.Wait()
}
