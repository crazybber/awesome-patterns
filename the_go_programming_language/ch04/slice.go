package main

import "fmt"

func main() {
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr"}
	// create a new slice by syntax [m:n]
	m1 := months[1:3]
	m2 := months[2:4]
	// Two slices point to the same underlaying array, when one changes the other changes also
	fmt.Println(m1, m2)
	m2[0] = "Feb updated"
	fmt.Print(m1, m2)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
