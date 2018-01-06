package main

import "github.com/jianhan/go-patterns/concurrency/subtasks/divide_and_conquer"

// https://medium.com/capital-one-developers/buffered-channels-in-go-what-are-they-good-for-43703871828

// One common pattern for goroutines is fan-out. When you want to apply the same data to multiple algorithms,
// you can launch a goroutine for each subtask, and then gather the data back in when they are done.
// For example, you might want to process the same data via multiple scoring algorithms and return back
// all of the scores or pull data from multiple microservices to compose a single page. A buffered channel is an
// ideal way to gather the data back from your subtasks.

func main() {
	divide_and_conquer.RunDivideAndConquer()
}
