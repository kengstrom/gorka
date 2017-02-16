package gorka

import (
	"math/rand"
	"fmt"
	//"sort"
	"sort"
)

const (
	HUMAN = iota
	ROBOT
)

var (
	ROBOT_NAMES []string = []string{"C-3PO", "R2-D2", "IG-88", "Optimus Prime", "Marvin", "T-800", "T-1000", "ED-209"}
)

type Player struct {
	Type int
	Name string
	Hand Cards
	Points int
}

// to string..
func (player Player) String() string {
	player.SortCards()
	if player.Type == ROBOT {
		return fmt.Sprintf("Robot player called %s (%s)", player.Name, player.Hand)
	} else if player.Type == HUMAN {
		return fmt.Sprintf("Human player called %s (%s)", player.Name, player.Hand)
	} else {
		return fmt.Sprintf("Alien player called %s (%s)", player.Name, player.Hand)
	}
}


func createHumanPlayer(name string) Player{
	var player = Player{HUMAN, name, nil, 0}
	return player
}

func createRobotPlayer() Player {
	var nameIndex = rand.Intn(len(ROBOT_NAMES))
	var name string = ROBOT_NAMES[nameIndex]
	ROBOT_NAMES = append(ROBOT_NAMES[:nameIndex], ROBOT_NAMES[nameIndex+1:]... )
	var player = Player{Type: ROBOT, Name: name}
	return player
}

func (player *Player) PlayCard(highCard Card) Card{
	var card Card
	if player.Type == ROBOT {
		if highCard.Value == 0 {
			fmt.Printf("%s is first to play.\n", player.Name)
		}
		// get playable cards...
		// find first playable card
		var firstPlayableCard = player.getPlayableCards(highCard)
		if firstPlayableCard == len(player.Hand) {
			// must play lowest card
			fmt.Printf("%s cannot play any card. Must play lowest\n", player.Name)
			card = player.Hand[0]
			player.Hand = append(player.Hand[:0], player.Hand[1:]...)
		} else {
			fmt.Printf("%s can play a card.\n", player.Name)
			var middleCard int
			middleCard = (len(player.Hand) - firstPlayableCard) / 2
			playedCard := rand.Intn(len(player.Hand) - middleCard - firstPlayableCard) + firstPlayableCard + middleCard
			card = player.Hand[playedCard]
			player.Hand = append(player.Hand[:playedCard], player.Hand[playedCard+1:]...)
		}
	}
	fmt.Printf("%s plays %s\n", player.Name, card)
	return card
}

func (player *Player) SortCards() {
	sort.Sort(player.Hand)
}

func (player *Player) getPlayableCards(highcard Card) int {
	for i := 0; i < len(player.Hand); i++ {
		if highcard.Value <= player.Hand[i].Value && highcard.Value != ACE {
			return i
		}
	}
	return len(player.Hand)
}

func (player *Player) receiveCard(card Card) {
	fmt.Print(player.Name)
	fmt.Print(" received ")
	fmt.Println(card)
	player.Hand = append(player.Hand, card)
}