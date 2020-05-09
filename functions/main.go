package main

import "github.com/davecgh/go-spew/spew"

type Cruncher func(int) int

func mul(n int) int {
	return n * 2
}

func add(n int) int {
	return n + 100
}

func sub(n int) int {
	return n - 1
}

func crunch(nums []int, a ...Cruncher) (rnums []int) {
	rnums = append(rnums, nums...)
	for _, f := range a {
		for i, n := range rnums {
			rnums[i] = f(n)
		}
	}
	return
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	spew.Dump(crunch(nums, mul, add, sub))
}
