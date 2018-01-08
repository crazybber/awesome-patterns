package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	barrier("http://httpbin.org/headers", "http://httpbin.org/User-Agen")
}

var timeoutMilliseconds int = 5000

type barrierResp struct {
	Err  error
	Resp string
}

func barrier(endpoints ...string) {
	requestNumber := len(endpoints)
	in := make(chan barrierResp, requestNumber)
	defer close(in)
	responses := make([]barrierResp, requestNumber)
	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}
	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
			break
		}
		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}
