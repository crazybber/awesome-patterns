package internal

import (
	"fmt"
	"sync"
)

// singleton is private struct, it should be created and fetched by GetSingletonObject func
type singleton struct {
}

func (singleton) SayHi() {
	fmt.Println("Hi!")
}

var (
	once     sync.Once
	instance singleton
)

func GetSingletonObject() singleton {
	once.Do(func() {
		instance = singleton{}
	})

	return instance
}
