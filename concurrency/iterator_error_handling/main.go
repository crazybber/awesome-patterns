package main

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

type FetchResult struct {
	Domain     string
	StatusCode uint
	Header     string
}

func dummyFetchUrl(url string) *FetchResult {
	time.Sleep(time.Second * 1)
	return &FetchResult{
		url,
		200,
		"Dummy Header",
	}
}

func main() {
	r := dummyFetchUrl("http://www.google.com")
	spew.Dump(r)
}
