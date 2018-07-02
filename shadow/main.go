package main

import "fmt"

func f() string {
	return "test"
}

func init() {
	fmt.Println(f())
}

var g = "g"

func main() {
	s := "hello world"
	p := &s
	*p = "H2"
	fmt.Println(s, s[0], len([]byte(s)))

}
