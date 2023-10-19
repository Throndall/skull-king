package commands

import (
	"fmt"

	"github.com/metalblueberry/skull-king/pkg/game"
	"github.com/metalblueberry/skull-king/pkg/skullking"
)

type PlayCard struct {
	PlayerName string
	Card       skullking.Card
}

func (c PlayCard) Execute(state *game.State) error {

	player := state.GetPlayer(c.PlayerName)
	if player == nil {
		return fmt.Errorf("player not found")
	}
	
	if !player.Hand.Contains(c.Card) {
		return fmt.Errorf("the card is not in the hand")
	}

	player.Hand.Remove(c.Card)

	state.CurrentRound().CurrentTrick().Play(skullking.Play{
		PlayerName: c.PlayerName,
		Card:       c.Card,
	})

	return nil
}
