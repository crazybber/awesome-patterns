package basic

import (
	"fmt"
	"testing"
)

func TestPrintAddress(t *testing.T) {
	var a int
	fmt.Printf("%T, %v, %p \n", a, a, &a)
	passByVariable(a)
	passByPointer(&a)
}

func passByVariable(a int) {
	fmt.Printf("%T, %v, %p \n", a, a, &a)
}

func passByPointer(a *int) {
	fmt.Printf("%T, %v, %p \n", a, a, &a)
	fmt.Printf("%T, %v, %p \n", *a, *a, &*a)
}

type robot struct{}

func TestStructAddress(t *testing.T) {
	var a robot
	fmt.Printf("%T, %v, %p \n", a, a, &a)
	passStructByVariable(a)
	passStructByPointer(&a)
}

func passStructByVariable(a robot) {
	fmt.Printf("[passStructByVariable] %T, %v, %p \n", a, a, &a)
}

func passStructByPointer(a *robot) {
	fmt.Printf("[passStructByPointer] %T, %v, %p \n", a, a, &a)
	fmt.Printf("[passStructByPointer] %T, %v, %p \n", *a, *a, &*a)
}
