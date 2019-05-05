// 外观模式 facade pattern.
// 为多个子模块提供一个统一的调用接口,子模块可以是同一接口的实现也可以不同.
// 实际上编写程序的时候很多地方不知不觉的会使用了模式.
// 这里以购买鸡蛋，牛奶，小麦粉为例，从代购处一次性购买三种而不需要分别访问三个商店
package main

type Shop interface {
	Sell()
}

type EggShop struct {}

func (EggShop)Sell(){
	println("no more eggs left")
}

type MilkShop struct {}

func (MilkShop)Sell(){
	println("no more milk left")
}

type WheatFlourShop struct {}

func (WheatFlourShop)Sell(){
	println("no more wheat flour left")
}

type DealerFacade struct {
	EgShop Shop
	MkShop Shop
	WfShop Shop
}

func (d DealerFacade)SellAll(){
	d.EgShop.Sell()
	d.MkShop.Sell()
	d.WfShop.Sell()
}

func main(){
	dealer := DealerFacade{
		EggShop{},
		MilkShop{},
		WheatFlourShop{},
	}
	dealer.SellAll()
}




