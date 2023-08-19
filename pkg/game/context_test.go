package game_test

import (
	"testing"

	"github.com/metalblueberry/skull-king/pkg/game"
	"github.com/metalblueberry/skull-king/pkg/game/commands"
)

func TestContext_NewGame(t *testing.T) {
	// Given
	context := game.NewContext()

	// When
	context.Push(commands.AddPlayer{
		PlayerName: "David",
	})

	// Then
	state := context.State()
	if state.Players[0].Name != "David" {
		t.Fatalf("Player in game is %s but wanted David", state.Players[0].Name)
	}

	// context.Commands[0].(commands.AddPlayer)

}
