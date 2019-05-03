// 状态模式 state pattern.
// 与策略模式在一些场景可以通用，但策略模式倾向于调用者根据情况手动改变内部策略以切换算法，
// 状态模式倾向于由Context内部自行管理状态，只需要设定初始状态即可，不需要手动切换.
//
// 以下是以跷跷板seesaw为例，分为[左侧高LeftState|右侧高RightState]
package main

import (
	"math/rand"
	"time"
)

// state
type SeesawState interface {
	LiftLeftSide(S *Seesaw)
	LiftRightSide(S *Seesaw)
}

type LeftState struct {}

func (LeftState) LiftLeftSide(S *Seesaw) {
	println("↑LEFT > left side wads already lifted")
}

func (l LeftState) LiftRightSide(S *Seesaw) {
	println("RIGHT↑ > lift right side")
	S.State = RightState{}
}

type RightState struct {}

func (r RightState) LiftLeftSide(S *Seesaw) {
	println("↑LEFT > lift left side")
	S.State = LeftState{}
}

func (RightState) LiftRightSide(S *Seesaw) {
	println("RIGHT↑ > right side wads already lifted")
}

// context
type Seesaw struct {
	State SeesawState
}

func (s *Seesaw)MakeLeftUp(){
	s.State.LiftLeftSide(s)
}

func (s *Seesaw)MakeRightUp(){
	s.State.LiftRightSide(s)
}

func main(){
	// init left
	seesaw := &Seesaw{
		State:LeftState{},
	}

	rand.Seed(time.Now().UnixNano())
	for i:=0 ;i<10 ;i++{
		if rand.Intn(2) == 1{
			// ▄▃▂
			seesaw.MakeLeftUp()
		}else{
			// ▂▃▄
			seesaw.MakeRightUp()
		}
	}
	/*
	output:
	RIGHT↑ > lift right side
	RIGHT↑ > right side wads already lifted
	RIGHT↑ > right side wads already lifted
	↑LEFT > lift left side
	↑LEFT > left side wads already lifted
	↑LEFT > left side wads already lifted
	↑LEFT > left side wads already lifted
	RIGHT↑ > lift right side
	RIGHT↑ > right side wads already lifted
	↑LEFT > lift left side
	 */
}

