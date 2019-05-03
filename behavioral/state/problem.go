// 存疑部分.
//
// 这里是以玩家的健康状态为基准，影响生命回复和受到伤害的例子，有三个状态[健康HealthyState|受伤WoundedState|死亡DeadState].
// 和main.go中的跷跷板例子由行为驱动状态变化不同.以下例子的行为是受到数据影响的,结果就是达不到完全消除if-else的效果
//
// 疑问：
// 1. 很多地方说状态模式可以消除if-else，事实上要做到自行的状态切换很多时候还是会用到if-else，这里的消除应当是指
// 将巨大的条件语句拆分开(?)
// 2. 还是说这种由数据决定状态的事务逻辑下不适合使用状态模式
package main

import (
	"errors"
	"fmt"
)

// state
type PlayerState interface {
	Heal(p *Player) error
	Hurt(p *Player ,dmg int)
}

type HealthyState struct {}

func (h HealthyState)Heal(p *Player) error{
	return nil
}

func (h HealthyState)Hurt(p *Player ,dmg int){
	if dmg > 0 && dmg < p.MaxHealth{
		p.Health = p.Health - dmg
		p.State = WoundedState{}
	}else if dmg > p.MaxHealth{
		p.Health = 0
		p.State = DeadState{}
	}
}

type WoundedState struct {}

func (WoundedState)Heal(p *Player) error{
	if p.Health >= p.MaxHealth - 5{
		fmt.Printf("healing from %d to %d\n" ,p.Health ,p.MaxHealth)
		p.State = HealthyState{}
		p.Health = p.MaxHealth
	}else{
		fmt.Printf("healing from %d to %d\n" ,p.Health ,p.Health+5)
		p.Health = p.Health + 5
	}
	return nil
}

func (h WoundedState)Hurt(p *Player ,dmg int){
	if p.Health > dmg{
		p.Health = p.Health - dmg
	}else {
		p.State = DeadState{}
		p.Health = 0
	}
}

type DeadState struct {}

func (DeadState)Heal(P *Player) error{
	return errors.New("you are dead")
}

func (DeadState)Hurt(P *Player ,dmg int){}

// context
type Player struct {
	Health int
	MaxHealth int
	State PlayerState
}

func (p *Player)HealPlayer()error{
	return p.State.Heal(p)
}

func (p *Player)HurtPlayer(damage int){
	fmt.Printf("damage %d\n" ,damage)
	p.State.Hurt(p ,damage)
}

//func main(){
//	player := &Player{
//		Health:100,
//		MaxHealth:100,
//		State:HealthyState{},
//	}
//
//	rand.Seed(time.Now().UnixNano())
//	for i:=0 ;i<10 ;i++{
//		if err := player.HealPlayer();err != nil{
//			fmt.Println(err)
//			break
//		}
//		player.HurtPlayer(rand.Intn(30))
//	}
//	/*
//	output:
//	damage 28
//	healing from 72 to 77
//	damage 29
//	healing from 48 to 53
//	damage 15
//	healing from 38 to 43
//	damage 1
//	healing from 42 to 47
//	damage 19
//	healing from 28 to 33
//	damage 27
//	healing from 6 to 11
//	damage 16
//	you are dead
//	 */
//}
