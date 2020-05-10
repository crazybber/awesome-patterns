package main

import "fmt"

type stackEntry struct {
	Next  *stackEntry
	Value interface{}
}

type stack struct {
	top *stackEntry
}

func (s *stack) Push(v interface{}) {
	var e stackEntry
	e.Value = v
	e.Next = s.top
	s.top = &e
}

func (s *stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	v := s.top.Value
	s.top = s.top.Next
	return v
}

func (s *stack) GetTop() *stackEntry {
	return s.top
}

func main() {
	s := &stack{}
	s.Push("one")
	s.Push("two")
	s.Push("three")
	fmt.Printf("%#v", s.GetTop().Next)
}
