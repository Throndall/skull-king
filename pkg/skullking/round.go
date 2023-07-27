package skullking

// Trick Consists of each player, in clockwise order, playing
// 1 card face-up to the table. The person who plays the
// highest valued card wins and takes the trick. The player
// gathers the cards in a stack before them.
type Trick struct {
	Table []*Play
}

// Round Consisting of 1 or more tricks. The number of
// tricks in a round is determined by the number of cards
// dealt. A round begins by dealing cards and ends when all
// cards dealt have been played.
type Round struct {
	Tricks []Trick
}

// Play is the Card played by a Player during a Trick
type Play struct {
	Player *Player
	Card   Card
}

// NewTrick creates a new empty Trick struct that can be used to Play cards by the Players
func NewTrick(numberOfPlayers int) *Trick {
	return &Trick{
		Table: make([]*Play, numberOfPlayers),
	}
}

// Play adds the Play action for the next player to the Trick
// It will return an error if the card cannot be played
func (t *Trick) Play(p Play) error {
	panic("Not Implemented Yet")
}

// Winner will return the player that wins the current Trick
func (t *Trick) Winner() *Player {

	actualWinner := t.Table[0]
	for i := 1; i < len(t.Table); i++ {
		if t.Table[i].Card.Type == CardTypeEscape {
			continue
		}

		if actualWinner.Card.Type == CardTypeEscape {
			actualWinner = t.Table[i]
			continue
		}

		if actualWinner.Card.Type == CardTypePirate {
			continue
		}

		if t.Table[i].Card.Type == CardTypePirate {
			actualWinner = t.Table[i]
			continue
		}

		if t.Table[i].Card.Type != actualWinner.Card.Type && t.Table[i].Card.Type != CardTypeSuitBlack {
			continue
		}

		if t.Table[i].Card.Type == CardTypeSuitBlack && actualWinner.Card.Type != CardTypeSuitBlack {
			actualWinner = t.Table[i]
			continue
		}

		if t.Table[i].Card.Value > actualWinner.Card.Value {
			actualWinner = t.Table[i]
			continue
		}
	}
	return actualWinner.Player
}

// Points returns the amount of points that this specific trick is worth for the player that wins it.
func (t *Trick) Points() int {
	return 0
}

func (t *Trick) Leading() CardType {

	position := 0
	for i := position; i < len(t.Table); i++ {
		if t.Table[i].Card.Type != CardTypeEscape {
			position = i
			break
		}
	}

	if t.Table[position].Card.Type != CardTypeSuitGreen &&
		t.Table[position].Card.Type != CardTypeSuitBlack &&
		t.Table[position].Card.Type != CardTypeSuitYellow &&
		t.Table[position].Card.Type != CardTypeSuitPurple {
		return CardTypeNone
	}

	return t.Table[position].Card.Type
}
