package main

import (
	"fmt"
	"time"
)

type Human interface {
	Speak()
}

type Australia struct {
}

func (h *Australia) Speak() {
	fmt.Println("I am Australia")
}

func InterfacePresentation(vs ...interface{}) {
	t := time.Now()
	for _, v := range vs {
		b := time.Since(t)
		fmt.Println("Before Checking", b)
		m, ok := interface{}(v).(Human)
		a := time.Since(t)
		fmt.Println("After Checking", a)
		if ok {
			m.Speak()
		}
	}

}

func Presentation(human ...Human) {
	for _, h := range human {
		h.Speak()
	}
}

func main() {
	var persons []interface{} = make([]interface{}, 100)
	for i := 0; i <= 100; i++ {
		persons = append(persons, &Australia{})
	}
	InterfacePresentation(persons...)
}
