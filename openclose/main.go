package main

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Legs() int {
	return 4
}

func (c Cat) PrintLegs() {
	fmt.Printf("I have %d legs", c.Legs())
}

type OctCat struct {
	Cat
}

func (o OctCat) Legs() int {
	return 5
}

func main() {
	var oct OctCat
	fmt.Println(oct.Legs())
	oct.PrintLegs()
}
