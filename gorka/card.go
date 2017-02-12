package gorka

import "fmt"

var (
	SUITES []string = []string{"C", "D", "H", "S"}
	VALUES []string = []string{"","","2", "3","4","5","6","7","8","9","10","J", "Q", "K", "A"}
//	SUITES []string = []string{"club", "diamond", "hearts", "spades"}
//	VALUES []string = []string{"","","two", "three","four","five","six","seven","eight","nine","ten","jack", "queen", "king", "ace"}
	/*
	GLYPH map[int]string = map[int]string {
		314:  "\U0001f0a1", 214:  "\U0001f0b1", 114:  "\U0001f0c1", 14:  "\U0001f0d1",
		302:  "\U0001f0a2", 202:  "\U0001f0b2", 102:  "\U0001f0c2", 2:   "\U0001f0d2",
		303:  "\U0001f0a3", 203:  "\U0001f0b3", 103:  "\U0001f0c3", 3:   "\U0001f0d3",
		304:  "\U0001f0a4", 204:  "\U0001f0b4", 104:  "\U0001f0c4", 4:   "\U0001f0d4",
		305:  "\U0001f0a5", 205:  "\U0001f0b5", 105:  "\U0001f0c5", 5:   "\U0001f0d5",
		306:  "\U0001f0a6", 206:  "\U0001f0b6", 106:  "\U0001f0c6", 6:   "\U0001f0d6",
		307:  "\U0001f0a7", 207:  "\U0001f0b7", 107:  "\U0001f0c7", 7:   "\U0001f0d7",
		308:  "\U0001f0a8", 208:  "\U0001f0b8", 108:  "\U0001f0c8", 8:   "\U0001f0d8",
		309:  "\U0001f0a9", 209:  "\U0001f0b9", 109:  "\U0001f0c9", 9:   "\U0001f0d9",
		310: "\U0001f0aa",  210: "\U0001f0ba",  110:  "\U0001f0ca", 10:  "\U0001f0da",
		311:  "\U0001f0ab", 211:  "\U0001f0bb", 111:  "\U0001f0cb", 11:  "\U0001f0db",
		312:  "\U0001f0ad", 212:  "\U0001f0bd", 112:  "\U0001f0cd", 12:  "\U0001f0dd",
		313:  "\U0001f0ae", 213:  "\U0001f0be", 113:  "\U0001f0ce", 13:  "\U0001f0de",
	}*/
	)
const (
	CLUBS = iota
/*	DIAMONDS
	HEARTS*/
	SPADES = 3
)
const (
	TWO = iota + 2
/*	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING*/
	ACE = 14
)

type Card struct {
	Suit	int
	Value	int
}
type Cards []Card

func (card Card)compareCard(card2 Card) bool {
	return false
}

// to string..
func (card Card) String() string {
	//return fmt.Sprintf("%s of %s", VALUES[card.Value], SUITES[card.Suit])
	return fmt.Sprintf("%s%s", VALUES[card.Value], SUITES[card.Suit])
	//return fmt.Sprintf("%s", GLYPH[card.Suit*100 + card.Value])
}

//implement sort.Interface
func (slice Cards) Len() int {
	return len(slice)
}

func (slice Cards) Less(i, j int) bool {
	return slice[i].Value < slice[j].Value
}

func (slice Cards) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}