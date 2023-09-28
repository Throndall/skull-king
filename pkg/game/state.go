package game

import "github.com/metalblueberry/skull-king/pkg/skullking"

// State represents the state of the game at a given moment
type State struct {
	Players []skullking.Player
	Deck    skullking.Deck
	Rounds  []skullking.Round
}

// PlayerCount returns the number of players in the game
func (s *State) PlayerCount() int {
	return len(s.Players)
}
