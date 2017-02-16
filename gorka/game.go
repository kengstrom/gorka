package gorka

import (
	"fmt"
)

type Game struct {
	Players[]  Player
	deck Deck
	PointLimit int
	Broker Broker
}

func (game *Game) addPlayer(player Player) {
	game.Players = append(game.Players, player)
}

func (game *Game) CreateGame(broker* Broker) {
	game.Broker = *broker
}

func (game *Game) CreateGame2(humans int, robots int) {
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
	/*
	game.deck = createDeck()
	//game.deck.printDeck()
	game.deck.shuffleDeck()
	//game.deck.printDeck()

	game.DealCards()
	*/
	//game.deck.printDeck()
	game.PointLimit = 30
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

func (game *Game) PlayDeal(startPlayer int) {
	var highestPlayer int = 0

	game.deck = createDeck()
	game.deck.shuffleDeck()
	game.clearHands()
	game.DealCards()

	for i:=0; i<6; i++ {
		fmt.Printf("Hand %d\n", i)
		fmt.Println("=================")
		var highCard Card
		highCard.Value = 0
		for j:=0; j<len(game.Players); j++ {
			var currentPlayer = (startPlayer + j) % len(game.Players)
			var playedCard = game.Players[currentPlayer].PlayCard(highCard)

			if playedCard.Value >= highCard.Value {
				highCard = playedCard
				highestPlayer = currentPlayer
			}
		}
		startPlayer = highestPlayer
	}

	var highestCard int = 0
	fmt.Println("Deal is over")
	for _, player := range game.Players {
		fmt.Printf("%s has a %s\n", player.Name, player.Hand[0])
		if highestCard < player.Hand[0].Value {
			highestCard = player.Hand[0].Value
		}
	}
	//for _, player := range game.Players {
	for i:=0; i<len(game.Players); i++ {
		if highestCard == game.Players[i].Hand[0].Value {
			game.Players[i].Points += highestCard
		}
	}
	game.printStandings()
}

func (game *Game) removePlayers() {
	for i:=len(game.Players)-1; i>=0; i-- {
		if game.Players[i].Points >= game.PointLimit {
			fmt.Printf("%s got %d points and is now out...\n", game.Players[i].Name, game.Players[i].Points)
			game.Players = append(game.Players[:i], game.Players[i+1:]...)
		}
	}
}

func (game *Game) clearHands() {
	for _, player := range game.Players {
		player.Hand = nil
	}
}
func (game *Game) PlayGame() {
	var startPlayer int = 0
	for len(game.Players) > 1 {
		game.PlayDeal(startPlayer)
		startPlayer ++
		game.removePlayers()
		if startPlayer > len(game.Players) {
			startPlayer = 0
		}
	}
	for _, player := range game.Players {
		fmt.Println(player)
	}
}
func (game *Game) printStandings() {
	for _, player := range game.Players {
		fmt.Printf("%s has %d points\n", player.Name, player.Points)
	}

}