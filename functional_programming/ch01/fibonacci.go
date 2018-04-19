package main

import "fmt"

type Memoized func(int) int

var fibMem = Memoize(fib)

func Memoize(mf Memoized) Memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

func FibMemoized(n int) int {
	return fibMem(n)
}

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return fib(x-2) + fib(x-1)
	}
}

func main() {
	fmt.Println(fib(30))
	fmt.Println(FibMemoized(30))
}
