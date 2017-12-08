package main

import (
	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	s := "TTTT"
	ps := &s
	spew.Dump("ORIGINAL PS", ps)
	spew.Dump("ADDRESS OF PS", &ps)

	pass(ps)
	spew.Dump("The Value Of S", s)
	p1 := &Person{"Jian", 23}
	passStruct(p1)
	spew.Dump("STRUCT AFTER PASSS STRUCT", p1)

}

func pass(ps *string) {
	*ps = "ZZZZZZZZZZZZZZZZZ"
	spew.Dump("PS IN PASS FUNCIONT", ps)
	spew.Dump("ADDRESS OF PS IN PASS FUNCIONT", &ps)
	ps = nil
}

func passStruct(p *Person) {
	p.Name = "Steve"
	p = nil
}

func passStruct2(p Person) {
	p.Name = "Steve"
}

func passVAL(v string) {
	v = ""
}
