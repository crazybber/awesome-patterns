package main

type GameType int

const (
	TypeFPS GameType = 1
	TypeRPG			 = TypeFPS << 1
)

type Game interface {
	Type() GameType
	Start(player string)
}

// chain of responsibility
type GameSelector struct {
	GameList []Game
}

func (g *GameSelector)AddGame(games ...Game){
	g.GameList = append(g.GameList ,games...)
}

func (g GameSelector) Start(t GameType, player string) {
	for _ ,v := range g.GameList{
		if v.Type() == t{
			v.Start(player)
			return
		}
	}
}

type FPSGame struct {
	t GameType
}

func (f FPSGame) Start(player string) {
	println(player ,"join in fps game")
}

func (f FPSGame)Type() GameType{
	return f.t
}

type RPGGame struct {
	t GameType
}

func (RPGGame) Start(player string) {
	println(player ,"join in rpg game")
}

func (r RPGGame)Type() GameType{
	return r.t
}

func main(){
	fps := FPSGame{TypeFPS}
	rpg := RPGGame{TypeRPG}

	sl := GameSelector{}
	sl.AddGame(fps ,rpg)

	player := "icg"
	sl.Start(TypeRPG ,player)
	println()
	sl.Start(TypeFPS ,player)
	// output:
	/*
	icg join in rpg game

	icg join in fps game
	 */
}
