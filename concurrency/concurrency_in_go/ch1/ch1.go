package ch1

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

func outputText(j *Job, goGroup *sync.WaitGroup) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
	goGroup.Done()
}
func Run() {
	goGroup := new(sync.WaitGroup)
	fmt.Println("Starting")
	hello := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 2
	world := new(Job)
	world.text = "world"
	world.i = 0
	world.max = 2
	go outputText(hello, goGroup)
	go outputText(world, goGroup)
	goGroup.Add(2)
	goGroup.Wait()
}
