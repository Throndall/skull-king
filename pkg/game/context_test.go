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
		PlayerName: "Erik",
	})
	context.Push(commands.AddPlayer{
		PlayerName: "Anna",
	})
	context.Push(commands.AddPlayer{
		PlayerName: "Ignacio",
	})
	context.Push(commands.AddPlayer{
		PlayerName: "Pablo",
	})
	context.Push(commands.AddPlayer{
		PlayerName: "Victor",
	})
	context.Push(commands.StartGame{})
	// Then
	state := context.State()
	if state.PlayerCount() != 5 {
		t.Fatalf("number of player incorrect %d,expected 5", state.PlayerCount())

	}
	for i := 0; i < len(state.Players); i++ {
		if len(state.Players[i].Hand) != 1 {
			t.Fatalf("player %s ,unexpected number card %d, expected 1 ", state.Players[i].Name, len(state.Players[i].Hand))
		}
	}

}
