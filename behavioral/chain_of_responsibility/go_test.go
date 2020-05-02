package main

import (
	"log"
	"sync"
	"testing"
	"time"
)

// 想弄清楚slice遍历时被修改到底会发生什么.
// 结果是安全失败.
// 并不意味着是线程安全的.
func TestGo(t *testing.T){
	s := make([]int ,0)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	for i:=0 ;i<10 ;i++{
		s = append(s, i)
	}

	itv := time.Second
	go func() {
		defer wg.Done()
		for k ,v := range s{
			println(k ,v)
			time.Sleep(itv)
		}
	}()

	time.Sleep(itv)

	go func() {
		defer wg.Done()
		s = append(s ,810)
		time.Sleep(itv)
		s = append(s ,114)
		log.Println("add")
	}()

	wg.Wait()
	log.Println(s)
}
