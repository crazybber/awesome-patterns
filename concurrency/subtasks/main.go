package main

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

type House struct {
	person *Person
	dog    *Dog
}

func main() {
	// divide_and_conquer.RunDivideAndConquer()
	// fetchers.RunFetchers()
	// h := &House{}
	// spew.Dump(h)

	// go func() {
	// 	defer fmt.Println("Returned")
	// 	for {
	// 		select {
	// 		default:
	// 			return
	// 		}
	// 	}
	// 	fmt.Println("TEST")
	// }()
	// time.Sleep(time.Second)
	// spew.Dump(caller())
	// spew.Dump(nakedReturn())
}

func nakedReturn() (i uint32, err error) {
	return 12, errors.New("Test")
}

func caller() (d *Dog, err error) {
	p, err := produceErr()
	spew.Dump(p)
	return
}

func produceErr() (*Person, error) {
	return &Person{Name: "James"}, errors.New("New Error")
}
