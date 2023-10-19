package commands

import (
	"github.com/metalblueberry/skull-king/pkg/game"
	"github.com/metalblueberry/skull-king/pkg/skullking"
)

type AddPlayer struct {
	PlayerName string
}

func (c AddPlayer) Execute(state *game.State) error {
	state.Players = append(state.Players, &skullking.Player{
		Name:        c.PlayerName,
		Scoresheets: []skullking.Scoresheet{},
	})
	return nil
}
