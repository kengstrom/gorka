package gorka

import "fmt"

type Game struct {
	Players[]  Player
}

func (game Game) CreateGame(humans int, robots int) {
	fmt.Println("Entering CreateGame")
}