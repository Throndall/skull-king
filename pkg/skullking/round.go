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

func (t *Trick) ContainsFigure() bool {
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.IsFigure() {
			return true
		}
	}
	return false
}

func (t *Trick) ContainsSkullKing() bool {
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.Type == CardTypeSkullKing {
			return true
		}
	}
	return false
}

func (t *Trick) ContainsMermaid() bool {
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.Type == CardTypeMermaid {
			return true
		}
	}
	return false
}

func (t *Trick) ContainsPirate() bool {
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.Type == CardTypePirate {
			return true
		}
	}
	return false
}

func (t *Trick) ContainsBlack() bool {
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.Type == CardTypeSuitBlack {
			return true
		}
	}
	return false
}

func (t *Trick) ContainsAllEscape() bool {
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.Type != CardTypeEscape {
			return false
		}
	}
	return true
}

func (t *Trick) NumberFigure() int {
	count := 0
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.IsFigure() {
			count++
		}

	}
	return count
}

func (t *Trick) FindMermaid() *Player {
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.Type == CardTypeMermaid {
			return t.Table[i].Player
		}
	}
	//panic("la sirena no ha sido encontrada")
	return nil
}

func (t *Trick) FindPirate() *Player {
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.Type == CardTypePirate {
			return t.Table[i].Player
		}
	}
	//panic("la sirena no ha sido encontrada")
	return nil
}

func (t *Trick) FindSkullKing() *Player {
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.Type == CardTypeSkullKing {
			return t.Table[i].Player
		}
	}
	//panic("la sirena no ha sido encontrada")
	return nil
}

func (t *Trick) FindHigestBlack() *Player {
	higest := &Play{}
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.Type != CardTypeSuitBlack {
			continue
		}
		if t.Table[i].Card.Value > higest.Card.Value {
			higest = t.Table[i]
		}
	}
	return higest.Player

}

func (t *Trick) FindHigestLeading() *Player {
	higest := &Play{}

	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.Type != t.Leading() {
			continue
		}
		if t.Table[i].Card.Value > higest.Card.Value {
			higest = t.Table[i]
		}

	}
	return higest.Player

}

func (t *Trick) FindFigure() *Player {
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.IsFigure() {
			return t.Table[i].Player
		}
	}
	//panic("No hemos encontrado la figura")
	return nil
}

func (t *Trick) FindDiferentFigures() int {
	if t.ContainsSkullKing() && t.ContainsMermaid() && t.ContainsPirate() {
		return 3
	}

	if (t.ContainsSkullKing() && t.ContainsMermaid()) || (t.ContainsSkullKing() && t.ContainsPirate()) || (t.ContainsMermaid() && t.ContainsPirate()) {
		return 2
	}

	if t.ContainsSkullKing() || t.ContainsMermaid() || t.ContainsPirate() {
		return 1
	}
	return 0
}

// Winner will return the player that wins the current Trick
func (t *Trick) Winner() *Player {
	if t.FindDiferentFigures() > 0 {
		if t.FindDiferentFigures() == 3 {
			return t.FindMermaid()
		}
		if t.FindDiferentFigures() == 1 {
			return t.FindFigure()
		}
		if t.FindDiferentFigures() == 2 {
			if t.ContainsSkullKing() && !t.ContainsMermaid() {
				return t.FindSkullKing()
			}
			if t.ContainsSkullKing() && t.ContainsMermaid() {
				return t.FindMermaid()
			}
			return t.FindPirate()
		}
	}
	if t.ContainsBlack() {
		return t.FindHigestBlack()
	}
	if t.ContainsAllEscape() {
		return t.Table[0].Player
	}
	return t.FindHigestLeading()
}

// Points returns the amount of points that this specific trick is worth for the player that wins it.
func (t *Trick) Bonus() int {
	Points := 0
	if t.FindDiferentFigures() > 0 {
		if t.FindDiferentFigures() == 3 {
			Points += 40
		}
		if t.FindDiferentFigures() == 2 {
			if t.ContainsPirate() && t.ContainsMermaid() {
				for i := 0; i < len(t.Table); i++ {
					if t.Table[i].Card.Type == CardTypeMermaid {
						Points += 20
					}
				}
			}
			if t.ContainsSkullKing() && t.ContainsPirate() {
				for i := 0; i < len(t.Table); i++ {
					if t.Table[i].Card.Type == CardTypePirate {
						Points += 30
					}
				}
			}
			if t.ContainsSkullKing() && t.ContainsMermaid() {
				Points += 40
			}
		}
	}
	for i := 0; i < len(t.Table); i++ {
		if t.Table[i].Card.Value == 14 {
			if t.Table[i].Card.Type == CardTypeSuitBlack {
				Points += 20
			} else {
				Points += 10
			}
		}
	}

	return Points
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
