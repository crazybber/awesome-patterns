package main

import "fmt"

// interface in Go provides both a value and pointer semantic form.
// An interface can store its own copy of a value (value semantics), or a value can be shared
// with the interface by storing a copy of the value’s address (pointer semantics).
// This is where the value/pointer semantics come in for interfaces

type printer interface {
	print()
}

type user struct {
	name string
}

func (u user) print() {
	fmt.Println("User Name:", u.name)
}

func main() {
	u := user{"Bill"}
	entities := []printer{
		u,
		&u,
	}
	u.name = "Bill_CHG"
	for _, e := range entities {
		e.print()
	}
}

func ptest(p printer) {
	p.print()
}
