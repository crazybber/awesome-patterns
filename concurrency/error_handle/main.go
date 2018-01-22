package main

// n general, your concurrent processes should send their
// errors to another part of your program that has complete information
// about the state of your program, and can make a more informed decision about what to do.
import (
	"fmt"
	"net/http"
)

func main() {
	// noErrorHandling()
	errorHandling()
}

// Here we create a type that encompasses both the *http.Response and the error
// possible from an iteration of the loop within our goroutine.
type Result struct {
	Error    error
	Response *http.Response
}

func errorHandling() {
	// This line returns a channel that can be read from to retrieve results of an iteration of our loop
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				// Here we create a Result instance with the Error and Response fields set.
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					// TODO: need to refactor the code, this might be a bug since
					// this will never execute.
					return
				// This is where we write the Result to our channel
				case results <- result:
				}
			}
		}()
		return results
	}
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		// Here, in our main goroutine, we are able to deal with errors coming out of the
		// goroutine started by checkStatus intelligently, and with the full context of the larger program.
		if result.Error != nil {
			fmt.Printf("error: %v", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}

func noErrorHandling() {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}
