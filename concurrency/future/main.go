package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type data struct {
	Body  []byte
	Error error
}

func futureData(url string) <-chan data {
	c := make(chan data, 1)

	go func() {
		var body []byte
		var err error

		resp, err := http.Get(url)
		defer resp.Body.Close()

		body, err = ioutil.ReadAll(resp.Body)

		c <- data{Body: body, Error: err}
	}()

	return c
}

func main() {
	future := futureData("https://medium.com/@thejasbabu/concurrency-patterns-golang-5c5e1bcd0833")
	future2 := futureData("https://golang.org/ref/mem")

	// do many other things

	body := <-future
	fmt.Printf("response: %#v", "test")
	fmt.Printf("error: %#v", body.Error)

	body2 := <-future2
	fmt.Printf("response: %#v", "test1")
	fmt.Printf("error: %#v", body2.Error)
}
