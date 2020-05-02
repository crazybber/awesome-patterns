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
	waitCh chan struct{}
	sign chan struct{}
}

func NewChanBarrier(gos int)*ChanBarrier{
	b := &ChanBarrier{
		waitCh:make(chan struct{} ,gos-1),
		sign:make(chan struct{}),
	}
	for i:=0;i<gos-1;i++{
		b.waitCh <- struct{}{}
	}
	return b
}

func (b *ChanBarrier) Await() error {
	select {
	case <- b.waitCh:
		<- b.sign
	default:
		close(b.sign)
		close(b.waitCh)
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
		go func(n int) {
			log.Println(n ,"await")
			if err := b.Await();err != nil{
				log.Println(err)
			}
			log.Println(n ,"pass")
			wg.Done()
		}(i)
	}
	wg.Wait()
}

