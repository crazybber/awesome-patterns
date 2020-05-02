// 外观模式 facade pattern.
// 为多个子模块提供一个统一的调用接口,子模块可以是同一接口的实现也可以不同.
// 实际上编写程序的时候很多地方不知不觉的会使用该模式.
// 这里以购买鸡蛋，牛奶，小麦粉为例，从代购处一次性购买三种食材而不需要分别访问三个商店,
// SellAll()方法还可以将三处接口包装为原子操作,在购买失败时进行回滚
package main

import (
	"errors"
	"fmt"
)

type Shop interface {
	Sell() error
}

type EggShop struct {}

func (EggShop)Sell() error{
	return errors.New("no more eggs left")
}

type MilkShop struct {}

func (MilkShop)Sell()error{
	return errors.New("no more milk left")
}

type WheatFlourShop struct {}

func (WheatFlourShop)Sell()error{
	return errors.New("no more wheat flour left")
}

type DealerFacade struct {
	EgShop Shop
	MkShop Shop
	WfShop Shop
}

func (d DealerFacade)BuyAll(){
	//if e := d.EgShop.Sell();e != nil{
	//	log.Println(e)
	//	RollBack()
	//}
	//...
	e1 := d.EgShop.Sell()
	e2 := d.MkShop.Sell()
	e3 := d.WfShop.Sell()
	if e1 == nil && e2 == nil && e3 == nil{
		//success
	}else{
		//fail and rollback
		fmt.Printf("error:\n%v\n%v\n%v" ,e1 ,e2 ,e3)
	}
}

func main(){
	dealer := DealerFacade{
		EggShop{},
		MilkShop{},
		WheatFlourShop{},
	}
	dealer.BuyAll()
	/*
	output:
	error:
	no more eggs left
	no more milk left
	no more wheat flour left
	 */
}




