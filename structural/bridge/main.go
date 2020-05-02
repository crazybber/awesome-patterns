// 桥接模式 bridge pattern.
// 桥接模式与策略模式有很多相似之处，原因在于他们都将行为交给了另一个接口实现，
// 桥接解决的是A+M、B+N、A+N等这样的调用者和行为的组合问题，是将不匹配的接口相组合,使之能够被统一的调用,一旦组合运行时一般不会发生变化.
// 策略模式则是为了动态的扩展，改变算法.本质上说桥接是结构模式策略是行为模式，侧重的分别是结构的组合和结构之间的通信
package main

// abstraction
type VehicleAbstraction interface {
	Run()
	MaxSpeed() int
}

// refined abstraction
type Truck struct {
	Driver DriverImplementor
	maxSpeed int
}

func (t Truck)Run(){
	println("truck begin to move...")
	t.Driver.Drive(t)
}

func (t Truck)MaxSpeed() int{
	return t.maxSpeed
}

// refined abstraction
type Car struct {
	Driver DriverImplementor
	maxSpeed int
}

func (c Car)Run(){
	println("car begin to move...")
	c.Driver.Drive(c)
}

func (c Car)MaxSpeed() int{
	return c.maxSpeed
}

// implementor
type DriverImplementor interface {
	Drive(vi VehicleAbstraction)
}

// concrete implementor
type RookieDriver struct {}

func (RookieDriver) Drive(vi VehicleAbstraction) {
	println("rookie driver -> speed" ,vi.MaxSpeed() >> 1 ,"km/h")
}

// concrete implementor
type OldDriver struct {}

func (OldDriver)Drive(vi VehicleAbstraction){
	println("old driver -> speed" ,vi.MaxSpeed() ,"km/h")
}

func main(){
	old := OldDriver{}
	rk := RookieDriver{}
	v1 := Truck{old ,60}
	v2 := Car{old ,130}
	v3 := Truck{rk ,60}
	v4 := Car{rk ,130}
	list := []VehicleAbstraction{v1 ,v2 ,v3 ,v4}
	for _ ,v := range list{
		v.Run()
	}
	/*
	output:
	truck begin to move...
	old driver -> speed 60 km/h
	car begin to move...
	old driver -> speed 130 km/h
	truck begin to move...
	rookie driver -> speed 30 km/h
	car begin to move...
	rookie driver -> speed 65 km/h
	 */

}






