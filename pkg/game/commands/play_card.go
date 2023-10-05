package commands

import (
	"github.com/metalblueberry/skull-king/pkg/game"
	"github.com/metalblueberry/skull-king/pkg/skullking"
)

type PlayCard struct {
	PlayerName string
	Card       skullking.Card
}

func (c PlayCard) Execute(state *game.State) error {

	round := state.Rounds[len(state.Rounds)-1]
	trick := round.Tricks[len(round.Tricks)-1]

	for i := 0; i < state.PlayerCount(); i++ {

		if state.Players[i].Name == c.PlayerName {
			trick.Table = append(trick.Table, &skullking.Play{
				Player: &state.Players[i],
				Card:   c.Card,
			})

			state.Players[i].Hand.Remove(c.Card)

			//state.Players[i].Hand
		}
	}

	return nil
}
