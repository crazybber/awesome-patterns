package main

import (
	"github.com/weichou1229/go-patterns/playground/singleton/internal"
)

func main() {
	var s = internal.GetSingletonObject()
	s.SayHi()
}
