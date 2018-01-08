package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// A Future indicates any data that is needed in future but its computation
// can be started in parallel so that it can be fetched from the background when needed.
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
	future := futureData("http://test.future.com")
	// do many other things
	body := <-future
	fmt.Printf("response: %#v", string(body.Body))
	fmt.Printf("error: %#v", body.Error)
}
