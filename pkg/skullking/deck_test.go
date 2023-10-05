package skullking

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeck_Draw(t *testing.T) {
	deck := NewDeck(1)
	deck.Put(
		Card{Number: 1},
		Card{Number: 2},
		Card{Number: 3},
		Card{Number: 4},
		Card{Number: 5},
		Card{Number: 6},
	)
	hand := deck.Draw(2)
	if len(hand) != 2 {
		t.Log("1")
		t.FailNow()
	}
	if hand[0].Number != 5 {
		t.Log("2")
		t.FailNow()
	}
	if hand[1].Number != 6 {
		t.FailNow()
	}
	if len(deck.Cards) != 4 {
		t.FailNow()
	}
	if deck.Cards[0].Number != 1 {
		t.FailNow()
	}
	if deck.Cards[1].Number != 2 {
		t.FailNow()
	}
	if deck.Cards[2].Number != 3 {
		t.FailNow()
	}
	if deck.Cards[3].Number != 4 {
		t.FailNow()
	}

}
func TestDeck_Draw_One(t *testing.T) {
	deck := NewDeck(1)
	deck.Put(
		Card{Number: 1},
		Card{Number: 2},
	)
	hand := deck.Draw(1)
	if len(hand) != 1 {
		t.Log("1")
		t.FailNow()
	}
	if hand[0].Number != 2 {
		t.Log("2")
		t.FailNow()
	}
	if deck.Cards[0].Number != 1 {
		t.FailNow()
	}

}

func TestDeck_Shufle(t *testing.T) {
	deck := NewDeck(2)
	deck.Put(
		Card{Number: 1},
		Card{Number: 2},
		Card{Number: 3},
		Card{Number: 4},
		Card{Number: 5},
		Card{Number: 6},
	)
	deck.Shufle()
	if deck.Cards[0].Number == 1 &&
		deck.Cards[1].Number == 2 &&
		deck.Cards[2].Number == 3 &&
		deck.Cards[3].Number == 4 &&
		deck.Cards[4].Number == 5 &&
		deck.Cards[5].Number == 6 {
		t.FailNow()
	}

}

func TestRemoveOneCard(t *testing.T) {

	cards := Cards{Card{Number: 2}}

	cards.Remove(Card{Number: 2})

	require.Empty(t, cards, "NO HAS QUITADO LA CARTA^^/")
}

func TestRemoveThree(t *testing.T) {

	cards := Cards{Card{Number: 2}, Card{Number: 1}, Card{Number: 3}}

	cards.Remove(Card{Number: 2})

	require.Len(t, cards, 2, "NO HAS QUITADO LA CARTA^^/")

	require.NotContains(t, cards, Card{Number: 2})

}
