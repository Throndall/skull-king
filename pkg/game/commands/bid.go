package commands

import (
	"github.com/metalblueberry/skull-king/pkg/game"
	"github.com/metalblueberry/skull-king/pkg/skullking"
)

type Bid struct {
	PlayerName string
	Bid        int
}

func (c Bid) Execute(state *game.State) error {
	for i := 0; i < len(state.Players); i++ {
		if state.Players[i].Name == c.PlayerName {
			state.Players[i].Scoresheets = append(state.Players[i].Scoresheets, skullking.Scoresheet{
				Bid: c.Bid,
			})
		}

	}

	return nil
}
