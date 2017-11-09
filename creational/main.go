package main

import (
	"fmt"

	"github.com/jianhan/go-patterns/creational/singleton"
)

func main() {
	instance := singleton.GetInstance()
	fmt.Println(instance.AddOne())
}
