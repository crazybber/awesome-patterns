package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job) {
	fileName := j.text + ".txt"
	fileContents := ""
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fileContents += j.text
		fmt.Println(j.text)
		j.i++
	}
	err := ioutil.WriteFile(fileName, []byte(fileContents), 0644)
	if err != nil {
		panic("Something went awry")
	}
}
func main() {
	hello := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 3
	world := new(Job)
	world.text = "world"
	world.i = 0
	world.max = 5
	go outputText(hello)
	go outputText(world)
	goSched()
}

func goSched() {
	iterations := 10
	for i := 0; i <= iterations; i++ {
		go showNumber(i)
	}
	runtime.Gosched()
	fmt.Println("Goodbye!")
}

func showNumber(num int) {
	fmt.Println(num)
}
