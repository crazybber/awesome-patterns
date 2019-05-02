// 主要分为四个部分，command ,ConcreteCommand ,receiver ,invoker.
// 分别对应以下GameCommand ,(CommandAttack|CommandEscape) ,GamePlayer ,Invoker.
package main

// command
type GameCommand interface {
	Execute()
}

// ConcreteCommand
type CommandAttack struct {Player GamePlayer}

func (c CommandAttack)Execute(){
	c.Player.Attack()
}

// ConcreteCommand
type CommandEscape struct {Player GamePlayer}

func (c CommandEscape)Execute(){
	c.Player.Escape()
}

// receiver
type GamePlayer interface {
	Attack()
	Escape()
}

type GunPlayer struct {Name string}

func (g GunPlayer) Attack() {
	println(g.Name ,"opened fire")
}

func (g GunPlayer) Escape() {
	println(g.Name ,"escape")
}

// invoker
type CommandInvoker struct {
	CommandList chan GameCommand
}

func (in *CommandInvoker)CallCommands(){
	for{
		select {
		case cmd := <- in.CommandList:
			cmd.Execute()
		default:
			return
		}
	}
}

func (in *CommandInvoker)PushCommands(c ...GameCommand){
	for _ ,v := range c{
		in.CommandList <- v
	}
}

func main(){
	invoker := &CommandInvoker{
		make(chan GameCommand ,10),
	}

	playerA := GunPlayer{"icg"}
	attk := CommandAttack{playerA}
	escp := CommandEscape{playerA}

	invoker.PushCommands(attk ,escp ,escp ,attk)
	invoker.CallCommands()
	// output:
	/*
		icg opened fire
		icg escape
		icg escape
		icg opened fire
	*/
}
