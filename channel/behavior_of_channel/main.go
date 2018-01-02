package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	ch := make(chan string)

	go func() {
		p := <-ch
		spew.Dump(p)
	}()

	ch <- "paper"

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	// fanout()
	// selectDrop()
	// waitForTasks()
	withTimeOut()
}

func fanout() {
	emps := 20
	// ch := make(chan string, emps)
	ch := make(chan string)

	for e := 0; e < emps; e++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
			ch <- "paper"
		}()
	}

	for emps > 0 {
		p := <-ch
		fmt.Printf("EMP %s %d \n", p, emps)
		emps--
	}
}

func selectDrop() {
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("employee : received :", p)
		}
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : send ack")
		default:
			fmt.Println("manager : drop")
		}
	}

	close(ch)
}

func waitForTasks() {
	ch := make(chan string, 1)
	defer close(ch)
	go func() {
		for p := range ch {
			fmt.Println("employee : working :", p)
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper"
	}
}

func withTimeOut() {
	duration := 50 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case p := <-ch:
		fmt.Println("work complete", p)

	case <-ctx.Done():
		fmt.Println("moving on")
	}
}
