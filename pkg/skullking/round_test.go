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
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: players[3],
		},
		{
			Name: "Different suit looses",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitParrot, Value: 14}},
				},
			},
			Want: players[2],
		},
		{
			Name: "Triumph wins against suit",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
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

type TableTestLeading struct {
	name  string
	trick Trick
	want  CardType
}

func TestLeading(t *testing.T) {
	table := []TableTestLeading{
		{
			name: "single suit",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeSuitGreen, Value: 2}},
					{Card: Card{Type: CardTypeSuitGreen, Value: 6}},
					{Card: Card{Type: CardTypeSuitGreen, Value: 7}},
				},
			},
			want: CardTypeSuitGreen,
		},
		{
			name: "three suit",
			trick: Trick{

				Table: []*Play{
					{Card: Card{Type: CardTypeSuitGreen}},
					{Card: Card{Type: CardTypeSuitYellow}},
					{Card: Card{Type: CardTypeSuitBlack}},
				},
			},
			want: CardTypeSuitGreen,
		},
		{
			name: "pirate",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypePirate}},
					{Card: Card{Type: CardTypeSuitYellow}},
					{Card: Card{Type: CardTypeSuitBlack}},
				},
			},
			want: CardTypeNone,
		},
	}
	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.trick.Leading()
			if actual != tt.want {
				t.Fatalf("the suit is diferent,actual: \"%s\",expected:\"%s\" ", actual, tt.want)
			}
		})
	}
}
