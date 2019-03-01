package syncmap

import (
	"fmt"
	"sync"
	"testing"
)

type Order struct {
	Id int
}

func TestSync(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)
	var cache sync.Map

	for i := 0; i < 10; i++ {
		go func() {
			cache.Store("test", &Order{})
			order, _ := cache.Load("test")
			o := order.(*Order)
			fmt.Println(o)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}

func TestSync2(t *testing.T) {
	var cache sync.Map

	cache.Store("test", nil)

	if order, ok := cache.Load("test"); ok && order != nil {
		fmt.Println(order)
	} else {
		fmt.Println("not exist")
	}
}

func TestSync3(t *testing.T) {
	var cache sync.Map
	orders := make([]Order, 2)

	cache.Store("test2", orders[0])
	cache.Store("test", orders[0])

	order, ok := cache.Load("test")
	fmt.Println(ok, order)

}
