package game

import "github.com/metalblueberry/skull-king/pkg/skullking"

// State represents the state of the game at a given moment
type State struct {
	Players []*skullking.Player
	Deck    skullking.Deck
	Rounds  []*skullking.Round
}

// PlayerCount returns the number of players in the game
func (s *State) PlayerCount() int {
	return len(s.Players)
}

// GetPlayer returns a player if exist, and nil if not
func (s *State) GetPlayer(name string) *skullking.Player {
	
	for i := 0; i < s.PlayerCount(); i++ {
		if s.Players[i].Name == name{
			return s.Players[i]
		}
	}
	return nil
}

// CurrentRound
func (s *State) CurrentRound() *skullking.Round{
	return s.Rounds[len(s.Rounds)-1]
}


