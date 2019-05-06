package main

import (
	"log"
	"sync"
)

type IBarrier interface {
	Await() error
}

type Barrier struct {
	cond *sync.Cond
	gos uint
	curgos uint
}

func NewBarrier(syncGos uint)*Barrier{
	if syncGos < 1{
		panic("min 1")
	}
	l := &sync.Mutex{}
	c := sync.NewCond(l)
	return &Barrier{
		cond:c,
		gos:syncGos,
	}
}

func (b *Barrier)Await() error{
	b.cond.L.Lock()
	defer b.cond.L.Unlock()
	b.curgos++
	if b.gos != b.curgos{
		b.cond.Wait()
	}else{
		b.curgos = 0
		b.cond.Broadcast()
	}
	return nil
}



func main(){
	var b IBarrier
	wg := &sync.WaitGroup{}
	gos := 10
	wg.Add(gos)
	b = NewBarrier(uint(gos))
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

	//b = NewTimeoutBarrier(time.Second ,2)
	//wg.Add(2)
	//
	//go func() {
	//	log.Println("await 1")
	//	if err := b.Await();err != nil{
	//		log.Println(err)
	//	}
	//	log.Println("pass 1")
	//	wg.Done()
	//}()
	//go func() {
	//	log.Println("await 2")
	//	// 导致前一个await超时
	//	time.Sleep(time.Second*4)
	//	if err := b.Await();err != nil{
	//		log.Println(err)
	//	}
	//	log.Println("pass 2")
	//	wg.Done()
	//}()
	//wg.Wait()
}

//// only use once
//type TimeoutBarrier struct {
//	ch chan struct{}
//	t time.Duration
//	gos uint
//	arrgos uint
//	togos int
//	m sync.Mutex
//}
//
//func NewTimeoutBarrier(t time.Duration ,syncgos uint)*TimeoutBarrier{
//	if t == 0{
//		t = math.MaxInt64
//	}
//	if syncgos < 1{
//		panic("min goroutine num = 1")
//	}
//	return &TimeoutBarrier{
//		ch:make(chan struct{}),
//		t:t,
//		gos:syncgos,
//	}
//}
//
//func (b *TimeoutBarrier)Await() (err error){
//	b.m.Lock()
//	b.arrgos++
//	if b.arrgos != b.gos{
//		b.m.Unlock()
//		select {
//		case <- b.ch:
//		case <- time.After(b.t):
//			b.togos++
//			err = errors.New("time out")
//		}
//	}else{
//		for i:=0;i<int(b.gos)-1-b.togos;i++{
//			b.ch <- struct{}{}
//		}
//		close(b.ch)
//		b.m.Unlock()
//	}
//	return err
//}
