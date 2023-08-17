package game_test

import (
	"testing"

	"github.com/metalblueberry/skull-king/pkg/game"
)

func TestContext_NewGame(t *testing.T) {
	context := game.NewContext()
	context.Execute(game.CommandAddPlayer{
		Player: game.Player{
			Name: "David",
		},
	})
	context.Commands[0].(game.CommandAddPlayer)
	
}
