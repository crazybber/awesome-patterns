package main

import "github.com/davecgh/go-spew/spew"

func main() {
	t1 := make(DogsList, 0)
	t2 := DogList{}
	spew.Dump(t1, t2)
}

type DogsList []*Dog

type Dog struct {
	Name string
}

func (d *Dog) GetName() string {
	return d.Name
}
