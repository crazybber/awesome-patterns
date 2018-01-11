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

func main() {
	nilFuncDefer()
}
