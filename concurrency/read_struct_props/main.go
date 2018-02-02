package main

import (
	"fmt"
	"time"
)

// package shows if read/write concurrently via struct

type Person struct {
	Name string
}

// simulate write operation for struct property which is name is this case
func (p *Person) updateName1() {
	for {
		p.Name = "Dummy Name 1"
	}
}

// simulate write operation for struct property which is name is this case
func (p *Person) updateName2() {
	for {
		p.Name = "Dummy Name 2"
	}
}

func (p *Person) printName() {
	for {
		fmt.Println("Current Name Is : ", p.Name)
	}
}

var m = map[string]int{"a": 1}

func main() {
	p := &Person{Name: "James"}
	go p.updateName1()
	go p.updateName2()
	go p.printName()
	// SimulateConcurrentReadWriteMap()
	time.Sleep(2 * time.Second)
}

func SimulateConcurrentReadWriteMap() {
	// Concurrent read is ok, but write is not
	go Read()
	go Write()
	time.Sleep(6 * time.Second)
}

func Read() {
	for {
		read()
	}
}

func Write() {
	for {
		write()
	}
}
func read() {
	_ = m["a"]
}

func write() {
	m["b"] = 2
}
