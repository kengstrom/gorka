package gorka

import "fmt"

type Game struct {
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
			game.Players[j].getCard(game.deck.dealCard())
		}
	}
}