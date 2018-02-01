package main

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

// package shows if read/write concurrently via struct

type Person struct {
	Name   string
	Age    uint32
	Hobbit string
}

func (p *Person) generateAge() {
	// mock read
	for {
		n := p.Name
		spew.Dump(n)
	}

}

func (p *Person) generateHobbit() {
	// mock read
	for {
		n := p.Name
		spew.Dump(n)
	}

}

func (p *Person) mockWrite1() {
	for {
		p.Name = "Nmae1"
	}
}

func (p *Person) mockWrite2() {
	for {
		p.Name = "Nmae2"
	}
}

var m = map[string]int{"a": 1}

func main() {
	p := &Person{Name: "James"}
	go p.generateAge()
	go p.generateHobbit()
	go p.mockWrite1()
	go p.mockWrite2()
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
