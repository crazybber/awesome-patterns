package main

import (
	"fmt"
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

func concurrentFetch(urls []string) <-chan *FetchResult {
	rChan := make(chan *FetchResult, len(urls))
	for _, url := range urls {
		go func(url string) {
			result := dummyFetchUrl(url)
			rChan <- result
		}(url)
	}
	return rChan
}

func main() {
	// mock the urls
	urls := []string{
		"test1",
		"test2",
		"test3",
		"test4",
		"test5",
		"test6",
	}

	// this is a prefered method
	chanResults := concurrentFetch(urls)
	for i := 0; i < len(urls); i++ {
		select {
		case c := <-chanResults:
			spew.Dump(c)
		}
	}

	// i := 0
	// for v := range concurrentFetch(urls) {
	// 	spew.Dump(i, v)
	// 	i++
	// 	if i == len(urls) {
	// 		break
	// 	}
	// }
	fmt.Println("All my things are done")
}
