// 中介者模式(mediator pattern)
// 当两个对象存在复杂的依赖关系时，考虑增加一个中介者使之解耦，使他们不需要进行显式的引用.
// 但问题在于中介本身会去实现复杂的逻辑，进而导致中介者变得复杂臃肿难以维护.
//
// 这里是竞价的例子.
// 参与竞价的玩家只需要向中介出价就可以知道是否能买到商品，由中介确定谁的出价高
package main

import (
	"log"
	"math/rand"
	"time"
)

// buyer 之间通过中介进行引用
type Buyer struct {
	m	*Mediator
	Name string// unique
}

func (b Buyer)Bid(price int){
	if p ,ok := b.m.isMaxPrice(b ,price);ok{
		log.Println(b.Name ,"are in lead temporarily with" ,p)
	}else{
		log.Println(b.Name ,"lose with" ,p)
	}
}

type Mediator struct {
	Buyers map[Buyer]int
	MaxPrice int
	MaxPriceBuyer Buyer
}

func (m *Mediator)isMaxPrice(b Buyer ,price int)(int ,bool){
	if b.Name == m.MaxPriceBuyer.Name{
		return m.MaxPrice ,true
	}
	m.Buyers[b] = price
	for k ,v := range m.Buyers{
		if v > m.MaxPrice{
			m.MaxPrice = v
			m.MaxPriceBuyer = k
		}
	}
	if b.Name == m.MaxPriceBuyer.Name{
		return price ,true
	}
	return price ,false
}

func main(){
	m := &Mediator{
		make(map[Buyer]int),
		0,
		Buyer{},
	}

	icg := Buyer{m ,"icg"}
	nyn := Buyer{m ,"nyn"}

	rand.Seed(time.Now().UnixNano())
	for i:=0 ;i<3 ;i++{
		icg.Bid(rand.Intn(10000))
		nyn.Bid(rand.Intn(10000))
	}
	// output:
	/*
	2019/05/02 12:19:52 icg are in lead temporarily with 1024
	2019/05/02 12:19:52 nyn are in lead temporarily with 4232
	2019/05/02 12:19:52 icg are in lead temporarily with 6412
	2019/05/02 12:19:52 nyn are in lead temporarily with 6747
	2019/05/02 12:19:52 icg lose with 4951
	2019/05/02 12:19:52 nyn are in lead temporarily with 6747
	 */
}
