package main

import "fmt"

type stackEntry struct {
	next  *stackEntry
	value interface{}
}

type stack struct {
	top *stackEntry
}

func (s *stack) Push(v interface{}) {
	var e stackEntry
	e.value = v
	e.next = s.top
	s.top = &e
}

func (s *stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	v := s.top.value
	s.top = s.top.next
	return v
}

func main() {
	s := &stack{}
	s.Push("one")
	s.Push("two")
	s.Push("three")
	fmt.Printf("%#v", s)
}
