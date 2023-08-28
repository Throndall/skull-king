package skullking

import "math/rand"

type Deck struct {
	Cards      []Card
	RandSource *rand.Rand
}

type Card struct {
	// Number identifies the card uniquely with a number
	Number int
	// Type of the card
	Type CardType
	// Value of the card if needed
	Value int
}

func (c Card) IsFigure() bool {
	if c.Type == CardTypeMermaid || c.Type == CardTypeSkullKing || c.Type == CardTypePirate {
		return true
	}
	return false
}

// que es un mazo un frenesi una sombra una ilusion

// cartas para todos de typo carta
// Los comentarios son terrorismo AT:Victor
// calla ya pesado
type Cards []Card

type CardValue int
type CardType string

const (
	CardTypeNone CardType = ""

	CardTypeSkullKing CardType = "SkullKing"
	CardTypeMermaid   CardType = "Mermaid"
	CardTypePirate    CardType = "Pirate"
	CardTypeEscape    CardType = "Escape"
	CardTypeTigress   CardType = "Tigress"

	CardTypeSuitParrot        CardType = "Parrot"
	CardTypeSuitGreen         CardType = "green"
	CardTypeSuitTreasureChest CardType = "Treasure Chest"
	CardTypeSuitYellow        CardType = "yellow"
	CardTypeSuitPirateMap     CardType = "Pirate Map"
	CardTypeSuitPurple        CardType = "purple"
	CardTypeSuitJollyRoger    CardType = "Jolly Roger"
	CardTypeSuitBlack         CardType = "black"
)

var (
	OriginalDeck = Deck{
		Cards: []Card{
			{Number: 1, Type: CardTypeSuitParrot, Value: 1},
			{Number: 2, Type: CardTypeSuitParrot, Value: 2},
			{Number: 3, Type: CardTypeSuitParrot, Value: 3},
			{Number: 4, Type: CardTypeSuitParrot, Value: 4},
			{Number: 5, Type: CardTypeSuitParrot, Value: 5},
			{Number: 6, Type: CardTypeSuitParrot, Value: 6},
			{Number: 7, Type: CardTypeSuitParrot, Value: 7},
			{Number: 8, Type: CardTypeSuitParrot, Value: 8},
			{Number: 9, Type: CardTypeSuitParrot, Value: 9},
			{Number: 10, Type: CardTypeSuitParrot, Value: 10},
			{Number: 11, Type: CardTypeSuitParrot, Value: 11},
			{Number: 12, Type: CardTypeSuitParrot, Value: 12},
			{Number: 13, Type: CardTypeSuitParrot, Value: 13},
			{Number: 14, Type: CardTypeSuitParrot, Value: 14},
			{Number: 15, Type: CardTypeSuitTreasureChest, Value: 1},
			{Number: 16, Type: CardTypeSuitTreasureChest, Value: 2},
			{Number: 17, Type: CardTypeSuitTreasureChest, Value: 3},
			{Number: 18, Type: CardTypeSuitTreasureChest, Value: 4},
			{Number: 19, Type: CardTypeSuitTreasureChest, Value: 5},
			{Number: 20, Type: CardTypeSuitTreasureChest, Value: 6},
			{Number: 21, Type: CardTypeSuitTreasureChest, Value: 7},
			{Number: 22, Type: CardTypeSuitTreasureChest, Value: 8},
			{Number: 23, Type: CardTypeSuitTreasureChest, Value: 9},
			{Number: 24, Type: CardTypeSuitTreasureChest, Value: 10},
			{Number: 25, Type: CardTypeSuitTreasureChest, Value: 11},
			{Number: 26, Type: CardTypeSuitTreasureChest, Value: 12},
			{Number: 27, Type: CardTypeSuitTreasureChest, Value: 13},
			{Number: 28, Type: CardTypeSuitTreasureChest, Value: 14},
			{Number: 29, Type: CardTypeSuitJollyRoger, Value: 1},
			{Number: 30, Type: CardTypeSuitJollyRoger, Value: 2},
			{Number: 31, Type: CardTypeSuitJollyRoger, Value: 3},
			{Number: 32, Type: CardTypeSuitJollyRoger, Value: 4},
			{Number: 33, Type: CardTypeSuitJollyRoger, Value: 5},
			{Number: 34, Type: CardTypeSuitJollyRoger, Value: 6},
			{Number: 35, Type: CardTypeSuitJollyRoger, Value: 7},
			{Number: 36, Type: CardTypeSuitJollyRoger, Value: 8},
			{Number: 37, Type: CardTypeSuitJollyRoger, Value: 9},
			{Number: 38, Type: CardTypeSuitJollyRoger, Value: 10},
			{Number: 39, Type: CardTypeSuitJollyRoger, Value: 11},
			{Number: 40, Type: CardTypeSuitJollyRoger, Value: 12},
			{Number: 41, Type: CardTypeSuitJollyRoger, Value: 13},
			{Number: 42, Type: CardTypeSuitJollyRoger, Value: 14},
			{Number: 43, Type: CardTypeSuitPirateMap, Value: 1},
			{Number: 44, Type: CardTypeSuitPirateMap, Value: 2},
			{Number: 45, Type: CardTypeSuitPirateMap, Value: 3},
			{Number: 46, Type: CardTypeSuitPirateMap, Value: 4},
			{Number: 47, Type: CardTypeSuitPirateMap, Value: 5},
			{Number: 48, Type: CardTypeSuitPirateMap, Value: 6},
			{Number: 49, Type: CardTypeSuitPirateMap, Value: 7},
			{Number: 50, Type: CardTypeSuitPirateMap, Value: 8},
			{Number: 51, Type: CardTypeSuitPirateMap, Value: 9},
			{Number: 52, Type: CardTypeSuitPirateMap, Value: 10},
			{Number: 53, Type: CardTypeSuitPirateMap, Value: 11},
			{Number: 54, Type: CardTypeSuitPirateMap, Value: 12},
			{Number: 55, Type: CardTypeSuitPirateMap, Value: 13},
			{Number: 56, Type: CardTypeSuitPirateMap, Value: 14},

			{Number: 57, Type: CardTypePirate},
			{Number: 58, Type: CardTypePirate},
			{Number: 59, Type: CardTypePirate},
			{Number: 60, Type: CardTypePirate},
			{Number: 61, Type: CardTypePirate},
			{Number: 62, Type: CardTypeTigress},
			{Number: 63, Type: CardTypeSkullKing},
			{Number: 64, Type: CardTypeMermaid},
			{Number: 65, Type: CardTypeMermaid},
			{Number: 66, Type: CardTypeEscape},
			{Number: 67, Type: CardTypeEscape},
			{Number: 68, Type: CardTypeEscape},
			{Number: 69, Type: CardTypeEscape},
			{Number: 70, Type: CardTypeEscape},
		},
	}
)

func NewDeck(seed int64) Deck {
	return Deck{
		Cards:      make([]Card, 0),
		RandSource: rand.New(rand.NewSource(seed)),
	}
}

func (d *Deck) Put(cards ...Card) {
	d.Cards = append(d.Cards, cards...)
}

func (d *Deck) Draw(n int) []Card {

	//d.Cards[len(d.Cards)-1]
	hand := d.Cards[len(d.Cards)-n:] //len(d.Cards)
	deck2 := d.Cards[:len(d.Cards)-n]
	d.Cards = deck2
	return hand
}

func (d *Deck) Shufle() {
	size := len(d.Cards)

	for i := 0; i < 7322; i++ {
		n1 := d.RandSource.Intn(size)
		n2 := d.RandSource.Intn(size)
		cardtem := d.Cards[n1]
		cardtem2 := d.Cards[n2]

		d.Cards[n1] = cardtem2
		d.Cards[n2] = cardtem
	}

}

func (d *Deck) Deal(hands int, cards int) []Cards {
	return nil
}
