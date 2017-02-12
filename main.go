package main

import (
	"github.com/kengstrom/gorka/gorka"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Starting game...")
	var game gorka.Game
	game.CreateGame(0,4)
	game.PlayGame()


}
