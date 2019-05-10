package main

import (
	"log"
	"sync"
)

type IBarrier interface {
	// error for timeout if need
	Await() error
}


// once
type ChanBarrier struct {
	gos int
	curGos int
	ch chan struct{}
	m sync.Mutex
}

func NewChanBarrier(gos int)*ChanBarrier{
	return &ChanBarrier{
		gos:gos,
		ch:make(chan struct{}),
	}
}

func (c *ChanBarrier) Await() error {
	c.m.Lock()
	if c.curGos++;c.gos != c.curGos{
		c.m.Unlock()
		<- c.ch
	}else{
		c.m.Unlock()
		close(c.ch)
	}
	return nil
}

//type Barrier struct {
//	cond *sync.Cond
//	gos uint
//	curgos uint
//}
//
//func NewBarrier(syncGos uint)*Barrier{
//	if syncGos < 1{
//		panic("min 1")
//	}
//	l := &sync.Mutex{}
//	c := sync.NewCond(l)
//	return &Barrier{
//		cond:c,
//		gos:syncGos,
//	}
//}
//
//func (b *Barrier)Await() error{
//	b.cond.L.Lock()
//	defer b.cond.L.Unlock()
//	b.curgos++
//	if b.gos != b.curgos{
//		b.cond.Wait()
//	}else{
//		b.curgos = 0
//		b.cond.Broadcast()
//	}
//	return nil
//}



func main(){
	var b IBarrier
	wg := &sync.WaitGroup{}
	gos := 10
	wg.Add(gos)
	//b = NewBarrier(uint(gos))
	b = NewChanBarrier(gos)
	for i:=0;i<gos ;i++ {
		go func() {
			log.Println("await")
			if err := b.Await();err != nil{
				log.Println(err)
			}
			log.Println("pass")
			wg.Done()
		}()
	}
	wg.Wait()
}

