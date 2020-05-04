package main

import (
	"github.com/crazybber/go-patterns/playground/singleton/internal"
)

func main() {
	var s = internal.GetSingletonObject()
	s.SayHi()
}
