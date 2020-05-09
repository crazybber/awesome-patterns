package main

import "fmt"

type user struct {
	name  string
	email string
}

func main() {
	users := []user{
		{"Bill", "bill@email.com"},
		{"Lisa", "lisa@email.com"},
		{"Nancy", "nancy@email.com"},
		{"Paul", "paul@email.com"},
	}

	for i, u := range users {
		fmt.Println(i, u)
	}
}
