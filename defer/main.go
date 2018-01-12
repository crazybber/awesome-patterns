package main

import "fmt"

// Sample code is taken from online tutorial https://blog.learngoprogramming.com/gotchas-of-defer-in-go-1-8d070894cb01
// it demostrate all the common GOTCHA when using defer.

// If a deferred func evaluates to nil, execution panic when surrounding func ends not when defer is called
func nilFuncDefer() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

// Do not use defer inside a loop unless you are sure about what you are doing. It may not work as expected.
// However, in some situations it will be handy for instance,delegating the recursivity of a func to a defer.
func deferInsideLoop() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	for _, v := range nums {
		defer func(v int) {
			fmt.Println(v)
		}(v)
	}
}

func main() {
	// nilFuncDefer()
	deferInsideLoop()
}
