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

// The deferred func above will only run when the func ends not when the deferred func’s surrounding block ends
// (the area inside curly braces containing the defer call). As seen in the example code, you can create separate blocks just using curly braces.
func block() {
	{
		defer func() {
			fmt.Println("block: defer runs")
		}()
	}
	fmt.Println("main: ends")
}

type Car struct {
	model string
}

func (c Car) PrintModel() {
	fmt.Println(c.model)
}

type Dog struct {
	name string
}

func (d *Dog) MakeSound() {
	fmt.Println("WOFF my name is", d.name)
}

// So, when a method with a value-receiver is used with defer, the receiver will be copied (in this case Car) at the time of
// registering and the changes to it wouldn’t be visible (Car.model). Because, the receiver is also an input param and evaluated immediately to “DeLorean DMC-12” when it’s
// registered with the defer.
func deferStruct() {
	c := Car{model: "model 1211"}
	defer c.PrintModel()
	c.model = "model 2018"
}

// Remember that the passed params to a deferred func are saved aside immediately without waiting the deferred func to be run.
func deferPointer() {
	d := &Dog{name: "James"}
	defer d.MakeSound()
	d.name = "Ann"
}

func main() {
	// nilFuncDefer()
	// deferInsideLoop()
	// block()
	// deferStruct()
	deferPointer()
}
