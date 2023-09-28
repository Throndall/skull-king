package commands

import (
	"testing"

	"github.com/metalblueberry/skull-king/pkg/game"
	"github.com/metalblueberry/skull-king/pkg/skullking"
)

func TestPlayCard(t *testing.T) {
	playCard := PlayCard{
		PlayerName: "Ortega el constructor",
		Card:       skullking.Card{Number: 25, Type: skullking.CardTypeSkullKing},
	}

	initialState := &game.State{
		Rounds: []skullking.Round{
			skullking.Round{
				Tricks: []skullking.Trick{skullking.Trick{
					Table: []*skullking.Play{},
				}},
			}},

		Players: []skullking.Player{
			skullking.Player{
				Name:        "Tximo",
				Scoresheets: []skullking.Scoresheet{},
				Hand:        skullking.Cards{skullking.Card{Number: 2, Type: skullking.CardTypeEscape}},
			},
			skullking.Player{
				Name:        "Anita",
				Scoresheets: []skullking.Scoresheet{},
				Hand:        skullking.Cards{skullking.Card{Number: 5, Type: skullking.CardTypeMermaid}},
			},
			skullking.Player{
				Name:        "Ortega el constructor",
				Scoresheets: []skullking.Scoresheet{},
				Hand:        skullking.Cards{skullking.Card{Number: 25, Type: skullking.CardTypeSkullKing}},
			},
		},
	}

	playCard.Execute(initialState)

	if len(initialState.Players[2].Hand) != 0 {
		t.Fatalf("Player %s not playing card", initialState.Players[2].Name)
	}

	if initialState.Rounds[0].Tricks[0].Table[0].Card.Number == 25 {
		t.Fatalf(" %s not playing card number 25", initialState.Players[2].Name)
	}

}
