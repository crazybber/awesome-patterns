// 享元模式 flyweight pattern.
// 通过保存不同特征的实例以达到复用的效果，从而节省内存和优化性能.
// 与对象池模式的不同之处在于享元模式所保存的实例具有不同的特征，而对象池则全部是相同的实例.
// 这里以蛋糕为例，字段flavour为其味道，根据不同的制作者将不同的实例保存起来.
// 当实例的特征较少时，享元模式还可以和单例模式相结合.
package main

type Cake struct {
	Flavour string
}

type CakeFactory struct {
	Cakes map[string]Cake
}

func (f CakeFactory)NewCake(flavour string)Cake{
	if c ,ok := f.Cakes[flavour];ok{
		println("get an existing" ,c.Flavour ,"cake from map")
	}else{
		f.Cakes[flavour] = Cake{flavour}
		println("put a new" ,flavour ,"cake into map")
	}
	return f.Cakes[flavour]
}

func main(){
	factory := CakeFactory{make(map[string]Cake)}

	factory.NewCake("strawberry")
	factory.NewCake("chocolates")
	factory.NewCake("nynicg")

	factory.NewCake("strawberry")
	factory.NewCake("nynicg")
	factory.NewCake("chocolates")
	/*
	output:
	put a new strawberry cake into map
	put a new chocolates cake into map
	put a new nynicg cake into map
	get an existing strawberry cake from map
	get an existing nynicg cake from map
	get an existing chocolates cake from map
	 */
}


