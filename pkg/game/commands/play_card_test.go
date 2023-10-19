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
		Rounds: []*skullking.Round{
			{
				Tricks: []*skullking.Trick{{
					Table: []*skullking.Play{},
				}},
			}},

		Players: []*skullking.Player{
			&skullking.Player{
				Name:        "Tximo",
				Scoresheets: []skullking.Scoresheet{},
				Hand:        skullking.Cards{skullking.Card{Number: 2, Type: skullking.CardTypeEscape}},
			},
			&skullking.Player{
				Name:        "Anita",
				Scoresheets: []skullking.Scoresheet{},
				Hand:        skullking.Cards{skullking.Card{Number: 5, Type: skullking.CardTypeMermaid}},
			},
			&skullking.Player{
				Name:        "Ortega el constructor",
				Scoresheets: []skullking.Scoresheet{},
				Hand:        skullking.Cards{skullking.Card{Number: 25, Type: skullking.CardTypeSkullKing}},
			},
		},
	}

	playCard.Execute(initialState)

	require.Len(t, initialState.Rounds[0].Tricks[0].Table, 1)
	for _, play := range initialState.Rounds[0].Tricks[0].Table {
		require.Equal(t, play.Card.Number, 25)
		require.Equal(t, play.PlayerName, "Ortega el constructor")
	}

	require.Emptyf(t, initialState.Players[2].Hand, "Player %s not playing card", initialState.Players[2].Name)

}

func TestPlayCard2(t *testing.T) {
	playCard := PlayCard{
		PlayerName: "Ortega el constructor",
		Card:       skullking.Card{Number: 29, Type: skullking.CardTypeSuitJollyRoger, Value: 1},
	}

	initialState := &game.State{
		Rounds: []*skullking.Round{
			{
				Tricks: []*skullking.Trick{{
					Table: []*skullking.Play{
						{
							PlayerName: "Tximo",
							Card:       skullking.Card{Number: 2, Type: skullking.CardTypeEscape},
						},
						{
							PlayerName: "Anita",
							Card:       skullking.Card{Number: 5, Type: skullking.CardTypeMermaid},
						},
					},
				}},
			}},

		Players: []*skullking.Player{
			&skullking.Player{
				Name:        "Tximo",
				Scoresheets: []skullking.Scoresheet{},
				Hand:        skullking.Cards{},
			},
			&skullking.Player{
				Name:        "Anita",
				Scoresheets: []skullking.Scoresheet{},
				Hand:        skullking.Cards{},
			},
			&skullking.Player{
				Name:        "Ortega el constructor",
				Scoresheets: []skullking.Scoresheet{},
				Hand:        skullking.Cards{skullking.Card{Number: 25, Type: skullking.CardTypeSkullKing}},
			},
		},
	}

	err := playCard.Execute(initialState)

	require.Error(t, err)

	require.Len(t, initialState.Players[2].Hand, 1)

	require.Len(t, initialState.Rounds[0].Tricks[0].Table, 2)
}


// This test prove that the first player in a trick is the same that the winner of the last one in the round.
func TestPlayCard_SecondTrick(t *testing.T) {
	playCard := PlayCard{
		PlayerName: "Anita",
		Card:       skullking.Card{Number: 6, Type: skullking.CardTypeMermaid},
	}

	initialState := &game.State{
		Rounds: []*skullking.Round{
			{
				Tricks: []*skullking.Trick{
					{

						Table: []*skullking.Play{
							{
								PlayerName: "Tximo",
								Card:       skullking.Card{Number: 2, Type: skullking.CardTypeEscape},
							},
							{
								PlayerName: "Anita",
								Card:       skullking.Card{Number: 5, Type: skullking.CardTypeMermaid},
							},
							{
								PlayerName: "Ortega el constructor",
								Card:       skullking.Card{Number: 25, Type: skullking.CardTypeSkullKing},
							},
						},
					},
				},
			},
			{
				Tricks: []*skullking.Trick{
					{
						Table: []*skullking.Play{
							{
								PlayerName: "Anita",
								Card:       skullking.Card{Number: 5, Type: skullking.CardTypeMermaid},
							},
							{
								PlayerName: "Ortega el constructor",
								Card:       skullking.Card{Number: 25, Type: skullking.CardTypeSkullKing},
							},
							{
								PlayerName: "Tximo",
								Card:       skullking.Card{Number: 2, Type: skullking.CardTypeEscape},
							},
						},
					},
					{
						Table: []*skullking.Play{},
					},
				},
			},

		},

		Players: []*skullking.Player{
			&skullking.Player{
				Name:        "Tximo",
				Scoresheets: []skullking.Scoresheet{},
				Hand:        skullking.Cards{skullking.Card{Number: 3, Type: skullking.CardTypeEscape}},
			},
			&skullking.Player{
				Name:        "Anita",
				Scoresheets: []skullking.Scoresheet{},
				Hand:        skullking.Cards{skullking.Card{Number: 6, Type: skullking.CardTypeMermaid}},
			},
			&skullking.Player{
				Name:        "Ortega el constructor",
				Scoresheets: []skullking.Scoresheet{},
				Hand:        skullking.Cards{skullking.Card{Number: 26, Type: skullking.CardTypeSkullKing}},
			},
		},
	}

	playCard.Execute(initialState)

	require.Len(t, initialState.Rounds[1].Tricks[1].Table, 1)
	for _, play := range initialState.Rounds[1].Tricks[1].Table {
		require.Equal(t, play.Card.Number, 6)
		require.Equal(t, play.PlayerName, "Anita")
	}

	require.Emptyf(t, initialState.Players[1].Hand, "Player %s not playing card", initialState.Players[1].Name)

}
