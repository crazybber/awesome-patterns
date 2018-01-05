package main

import "github.com/davecgh/go-spew/spew"

// https://medium.com/capital-one-developers/buffered-channels-in-go-what-are-they-good-for-43703871828

// One common pattern for goroutines is fan-out. When you want to apply the same data to multiple algorithms,
// you can launch a goroutine for each subtask, and then gather the data back in when they are done.
// For example, you might want to process the same data via multiple scoring algorithms and return back
// all of the scores or pull data from multiple microservices to compose a single page. A buffered channel is an
// ideal way to gather the data back from your subtasks.

func main() {
	evaluators := []Evaluator{google, yahoo, bing}
	data := "Query"
	r, e := DivideAndConquer(data, evaluators)
	spew.Dump(r, e)
}

var google = func(data interface{}) (interface{}, error) {
	return data, nil
}

var yahoo = func(data interface{}) (interface{}, error) {
	return data, nil
}

var bing = func(data interface{}) (interface{}, error) {
	return data, nil
}

type Evaluator func(data interface{}) (interface{}, error)

func DivideAndConquer(data interface{}, evaluators []Evaluator) ([]interface{}, []error) {
	gather := make(chan interface{}, len(evaluators))
	errors := make(chan error, len(evaluators))
	for _, v := range evaluators {
		go func(e Evaluator) {
			result, err := e(data)
			if err != nil {
				errors <- err
			} else {
				gather <- result
			}
		}(v)
	}
	out := make([]interface{}, 0, len(evaluators))
	errs := make([]error, 0, len(evaluators))
	for range evaluators {
		select {
		case r := <-gather:
			out = append(out, r)
		case e := <-errors:
			errs = append(errs, e)
		}
	}
	return out, errs
}
