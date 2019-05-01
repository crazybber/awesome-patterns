// 存疑的部分.
// 模板模式和策略模式在很多时候给人感觉基本可以划等号，事实上两者侧重的方向不同.
// 策略模式倾向于通过改变结构的属性从而改变算法，模板模式则倾向于事先定义一个处理的流程，该流程是可以被整体替换的.
// 个人认为应该保证模板模式下的结构体应当独立出来，每个流程单独被调用.
// 感觉以下两种写法都有点别扭，待指正

package main

type Game interface {
	OnStart()
	OnEnd()
}

type GameRunner struct {}

func (GameRunner)Go(g Game){
	g.OnStart()
	g.OnEnd()
}

type FPSGame struct {}

func (f FPSGame)OnStart(){
	println("fps game start")
}

func (f FPSGame)OnEnd(){
	println("fps game end")
}

type RPGGame struct {}

func (RPGGame) OnStart() {
	println("rpg game start")
}

func (RPGGame) OnEnd() {
	println("rpg game end")
}

func main(){
	tpl := GameRunner{}
	tpl.Go(FPSGame{})
	tpl.Go(RPGGame{})
	// output:
	/*
		fps game start
		fps game end
		rpg game start
		rpg game end
	 */
}


// ______________________________Another_______________

//type Game struct {
//	OnStart func()
//	OnEnd func()
//}
//
//func (g Game)Go(){
//	g.OnStart()
//	g.OnEnd()
//}
//
//type FPSGame struct {Game}
//
//func NewFPSGame()FPSGame{
//	g := Game{
//		OnStart: func() {
//			fmt.Println("start fps")
//			time.Sleep(time.Second / 2)
//		},
//		OnEnd: func() {
//			fmt.Println("you are killed")
//		},
//	}
//	return FPSGame{g}
//}
//
//type RPGGame struct {Game}
//
//func NewRPGGame()RPGGame{
//	g := Game{
//		OnStart: func() {
//			fmt.Println("start rpg")
//			time.Sleep(time.Second / 2)
//		},
//		OnEnd: func() {
//			fmt.Println("you win")
//		},
//	}
//	return RPGGame{g}
//}
//
//func main(){
//	NewFPSGame().Go()
//	println()
//	NewRPGGame().Go()
//}

