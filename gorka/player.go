package gorka

import (
	"math/rand"
	"fmt"
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
	Hand []Card
}

// to string..
func (player Player) String() string {
	if player.Type == ROBOT {
		return fmt.Sprintf("Robot player called %s (%s)", player.Name, player.Hand)
	} else if player.Type == HUMAN {
		return fmt.Sprintf("Human player called %s (%s)", player.Name, player.Hand)
	} else {
		return fmt.Sprintf("Alien player called %s (%s)", player.Name, player.Hand)
	}
}


func createHumanPlayer(name string) Player{
	var player = Player{HUMAN, name, nil}
	return player
}

func createRobotPlayer() Player {
	var nameIndex = rand.Intn(len(ROBOT_NAMES))
	var name string = ROBOT_NAMES[nameIndex]
	ROBOT_NAMES = append(ROBOT_NAMES[:nameIndex], ROBOT_NAMES[nameIndex+1:]... )
	var player = Player{ROBOT, name, nil}
	return player
}

func (player *Player) playCard(game *Game) {

}

func (player *Player) getCard(card Card) {
	fmt.Println(card)
	player.Hand = append(player.Hand, card)
}