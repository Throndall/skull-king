package skullking

import (
	"reflect"
	"testing"
)

type TableTestWinner struct {
	Name  string
	Trick Trick

	Want *Player
}

func TestTrick_Winner(t *testing.T) {
	players := []*Player{
		{Name: "Victor"},
		{Name: "Erik"},
		{Name: "Ortega"},
		{Name: "Ignacio"},
	}

	table := []TableTestWinner{
		{
			Name: "higher card wins the trick",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitJollyRoger, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitJollyRoger, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitJollyRoger, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitJollyRoger, Value: 4}},
				},
			},
			Want: players[3],
		},
		{
			Name: "Different suit looses",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitJollyRoger, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitJollyRoger, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitJollyRoger, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitParrot, Value: 14}},
				},
			},
			Want: players[2],
		},
		{
			Name: "Triumph wins against suit",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitJollyRoger, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitJollyRoger, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitJollyRoger, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitPirateMap, Value: 1}},
				},
			},
			Want: players[3],
		},
	}
	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			got := tt.Trick.Winner()
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("Trick.Winner() = %v, want %v", got, tt.Want)
			}
		})
	}
}
