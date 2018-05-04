package main

import (
	"fmt"
	"net/http"
	"time"

	"sync"

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
			wg := sync.WaitGroup{}
			wg.Add(len(urls))
			for _, url := range urls {
				go func(url string) {
					defer wg.Done()
					select {
					case <-done:
						return
					default:
						println("EXECUTE : ", url)
						resp, err := http.Get(url)
						results <- Result{Error: err, Response: resp}
					}
				}(url)
			}
			wg.Wait()
		}()
		return results
	}
	done := make(chan interface{}, 1)
	defer close(done)
	urls := []string{
		"https://www.google.com",
		"https://blog.golang.org/pipelines",
		"https://github.com/jianhan",
		"https://insights.stackoverflow.com/survey/2017",
		"https://hackernoon.com/top-10-python-web-frameworks-to-learn-in-2018-b2ebab969d1a",
		"https://blog.kowalczyk.info/article/1Bkr/3-ways-to-iterate-in-go.html",
	}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			done <- true
			fmt.Printf("error: %v\n", result.Error)
			break
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
	fmt.Println("End")
}
