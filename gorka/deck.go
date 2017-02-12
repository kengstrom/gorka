package gorka

import (
	"math/rand"
	"fmt"
)

type Deck struct {
	Cards Cards
}



func createDeck() Deck {
	var deck Deck
	for suit := CLUBS; suit <= SPADES; suit++ {
		for value := TWO; value <= ACE; value++ {
			card := Card{suit, value}
			deck.Cards = append(deck.Cards, card)
		}
	}
	fmt.Printf("New deck: %d cards in deck\n", len(deck.Cards) )
	return deck
}


func (deck *Deck) shuffleDeck() {
	for i := range deck.Cards {
		j := rand.Intn(i + 1)
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	}
}

func (deck *Deck) dealCard() Card {
	//var card  = deck.Cards[len(deck.Cards)-1]
	//var card = deck.Cards[0]
	//deck.Cards
	var card Card
	fmt.Printf("%d cards left in deck\n", len(deck.Cards) )
	card, deck.Cards = deck.Cards[len(deck.Cards)-1], deck.Cards[:len(deck.Cards)-1]
	//deck.Cards = append(deck.Cards[:0], deck.Cards[1:]...)
	return card
}

func (deck Deck)printDeck() {
	fmt.Println("=============")
	fmt.Println("Printing deck")
	fmt.Printf("%d cards in deck\n", len(deck.Cards))
	fmt.Println("=============")
	for _, card := range deck.Cards {
		fmt.Println(card)
	}
}