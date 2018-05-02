package main

import (
	"fmt"
	"net/http"
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
	v2()
	return
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

type Result struct {
	Error    error
	Response *http.Response
}

func v2() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {

				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}
	done := make(chan interface{})
	errCount := 0
	urls := []string{"a", "https://www.google.com", "b", "c", "d"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errCount++
			if errCount >= 3 {
				fmt.Println("Too many errors, breaking!")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
