package commands

import (
	"testing"

	"github.com/metalblueberry/skull-king/pkg/game"
	"github.com/metalblueberry/skull-king/pkg/skullking"
)

func TestBid(t *testing.T) {
	bid := Bid{
		PlayerName: "Orte",
		Bid:        1,
	}

	initialState := &game.State{
		Players: []skullking.Player{
			skullking.Player{
				Name:        "Tximo",
				Scoresheets: []skullking.Scoresheet{},
			},
			skullking.Player{
				Name:        "Anita",
				Scoresheets: []skullking.Scoresheet{},
			},
			skullking.Player{
				Name:        "Orte",
				Scoresheets: []skullking.Scoresheet{},
			},
		},
	}
	bid.Execute(initialState)

	if initialState.Players[2].Scoresheets[0].Bid != 1 {
		t.Fatalf("Player %s should have 1 bid and has %d bid", initialState.Players[2].Name, initialState.Players[2].Scoresheets[0].Bid)
	}
	if len(initialState.Players[0].Scoresheets) != 0 {
		t.Fatalf("Player %s should have 0 bid and has bid", initialState.Players[0].Name)
	}
	if len(initialState.Players[1].Scoresheets) != 0 {
		t.Fatalf("Player %s should have 0 bid and has bid", initialState.Players[1].Name)
	}

}

func TestBidRoundTwo(t *testing.T) {
	bid := Bid{
		PlayerName: "Anita",
		Bid:        2,
	}

	initialState := &game.State{
		Players: []skullking.Player{
			{
				Name: "Tximo",
				Scoresheets: []skullking.Scoresheet{
					{Bid: 0},
				},
			},
			{
				Name: "Anita",
				Scoresheets: []skullking.Scoresheet{
					{Bid: 0},
				},
			},
			{
				Name: "Orte",
				Scoresheets: []skullking.Scoresheet{
					{Bid: 1},
				},
			},
		},
	}
	bid.Execute(initialState)

	if len(initialState.Players[0].Scoresheets) != 1 {
		t.Fatalf("Player %s should have 0 bid and has bid", initialState.Players[0].Name)
	}
	if initialState.Players[1].Scoresheets[1].Bid != 2 {
		t.Fatalf("Player %s should have 0 bid and has %d bid", initialState.Players[1].Name, initialState.Players[2].Scoresheets[0].Bid)
	}
	if len(initialState.Players[2].Scoresheets) != 1 {
		t.Fatalf("Player %s should have 1 bid and has bid", initialState.Players[2].Name)
	}
}
