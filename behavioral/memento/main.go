// 备忘录模式 memento pattern
// 在不影响原结构封装的情况下，能够暂时保存一个结构的状态，并能够恢复
// 这里是一个游戏存档的例子，尝试保存玩家当前位置，并在读档的时候恢复
package main

import (
	"container/list"
	"log"
)

// originator
type Player struct {
	// 需要记录的数据可以考虑单独封装
	// type Pos struct{X,Y int}
	X,Y int

	// other info
	Name string
}

func (p *Player)MoveTo(x,y int){
	p.X = x
	p.Y = y
}

func (p Player)Save()PlayerMemento{
	return PlayerMemento{
		X:p.X,
		Y:p.Y,
	}
}

func (p *Player)Restore(m PlayerMemento){
	p.X = m.X
	p.Y = m.Y
}

// memento
type PlayerMemento struct {
	X,Y int
}

// caretaker
type PlayerCareTaker struct {
	MementoList *list.List
}

func (ct *PlayerCareTaker)AddMemento(memento PlayerMemento){
	ct.MementoList.PushFront(memento)
}

func (ct *PlayerCareTaker)RemoveLast()PlayerMemento{
	ele := ct.MementoList.Front()
	val := ct.MementoList.Remove(ele)
	if memento ,ok := val.(PlayerMemento);ok{
		return memento
	}else{
		return PlayerMemento{}
	}
}

func main(){
	ct := &PlayerCareTaker{list.New()}
	icg := &Player{
		X:114,
		Y:514,
		Name:"icg",
	}
	ct.AddMemento(icg.Save())
	log.Println(icg.X ,icg.Y)

	icg.MoveTo(810 ,19)
	log.Println(icg.X ,icg.Y)
	ct.AddMemento(icg.Save())

	icg.MoveTo(0 ,0)
	log.Println(icg.X ,icg.Y)

	icg.Restore(ct.RemoveLast())
	log.Println(icg.X ,icg.Y)

	icg.Restore(ct.RemoveLast())
	log.Println(icg.X ,icg.Y)
	/*
	output:
	2019/05/02 18:18:03 114 514
	2019/05/02 18:18:03 810 19
	2019/05/02 18:18:03 0 0
	2019/05/02 18:18:03 810 19
	2019/05/02 18:18:03 114 514
	 */
}
