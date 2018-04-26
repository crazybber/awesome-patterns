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
