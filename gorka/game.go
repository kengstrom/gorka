package gorka

import (
	"fmt"
)

type Game struct {
	NbrOfPlayers int
	Players[]  Player
	deck Deck
}

func (game *Game) addPlayer(player Player) {
	game.Players = append(game.Players, player)
}

func (game *Game) CreateGame(humans int, robots int) {
	fmt.Println("Entering CreateGame")
	for i:= 0 ; i<humans; i++ {
		game.addPlayer(createHumanPlayer("Player 1"))

	}
	for j:=0; j<robots; j++{
		game.addPlayer(createRobotPlayer())

	}
	game.NbrOfPlayers = humans + robots
	for _, player := range game.Players {
		fmt.Println(player)
	}
	game.deck = createDeck()
	//game.deck.printDeck()
	game.deck.shuffleDeck()
	//game.deck.printDeck()

	game.DealCards()

	//game.deck.printDeck()

	for _, player := range game.Players {
		fmt.Println(player)
	}

}

func (game *Game) DealCards() {
	for i:=0; i<7; i++ {
		for j := 0; j < len(game.Players); j++ {
			game.Players[j].receiveCard(game.deck.dealCard())
		}

	}
}


func (game *Game) PlayGame() {
	var startPlayer int = 0
	var highestPlayer int = 0
	for i:=0; i<6; i++ {
		fmt.Printf("Hand %d\n", i)
		fmt.Println("=================")
		var highCard Card
		highCard.Value = 0
		for j:=0; j<game.NbrOfPlayers; j++ {
			var currentPlayer = (startPlayer + j) % game.NbrOfPlayers
			var playedCard = game.Players[currentPlayer].PlayCard(highCard)

			if playedCard.Value >= highCard.Value {
				highCard = playedCard
				highestPlayer = currentPlayer
			}
		}
		startPlayer = highestPlayer
	}

	for _, player := range game.Players {
		fmt.Println(player)
	}

}