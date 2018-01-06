package main

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	a := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		test := "string"
		a <- test
	}()

	select {
	case t := <-a:
		spew.Dump(t)

	}

}
