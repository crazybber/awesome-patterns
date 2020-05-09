package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
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
	runtime.GOMAXPROCS(2)
	fmt.Printf("%d thread(s) available to Go.", listThreads())
}

func goSched() {
	iterations := 10
	for i := 0; i <= iterations; i++ {
		showNumber(i)
	}
	fmt.Println("Goodbye!")
}

func showNumber(num int) {
	tstamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(num, tstamp)
}

func listThreads() int {
	threads := runtime.GOMAXPROCS(0)
	return threads
}
