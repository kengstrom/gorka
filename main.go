package main

import (
	"github.com/kengstrom/gorka/gorka"
	"fmt"
)

func main() {
	fmt.Println("Starting game...")
	var game gorka.Game
	game.CreateGame(1,3)


}
