package commands

import (
	"testing"

	"github.com/metalblueberry/skull-king/pkg/game"
	"github.com/metalblueberry/skull-king/pkg/skullking"
	"github.com/stretchr/testify/require"
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

	require.Contains(t, initialState.Rounds[0].Tricks[0].Table, skullking.Card{Number: 25, Type: skullking.CardTypeSkullKing}, "not playing card number 25")

	require.Emptyf(t, initialState.Players[2].Hand, "Player %s not playing card", initialState.Players[2].Name)

}
