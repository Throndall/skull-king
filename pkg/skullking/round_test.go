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
		{Name: "0"},
		{Name: "1"},
		{Name: "2"},
		{Name: "3"},
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
					{Player: players[3], Card: Card{Type: CardTypeSuitGreen, Value: 14}},
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
		{
			Name: "Highest triumph wins",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitBlack, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitBlack, Value: 12}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 10}},
				},
			},
			Want: players[1],
		},
		{
			Name: "Pirate wins",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 12}},
					{Player: players[2], Card: Card{Type: CardTypePirate}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 10}},
				},
			},
			Want: players[2],
		},
		{
			Name: "Second pirate looses",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Player: players[1], Card: Card{Type: CardTypePirate}},
					{Player: players[2], Card: Card{Type: CardTypePirate}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 10}},
				},
			},
			Want: players[1],
		},
		{
			Name: "Escape looses",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeEscape}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitGreen, Value: 14}},
				},
			},
			Want: players[2],
		},
		{
			Name: "Escape everywhere",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeEscape}},
					{Player: players[1], Card: Card{Type: CardTypeEscape}},
					{Player: players[2], Card: Card{Type: CardTypeEscape}},
					{Player: players[3], Card: Card{Type: CardTypeEscape}},
				},
			},
			Want: players[0],
		},
		{
			Name: "Escape looses black",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeEscape}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			Want: players[3],
		},
		{
			Name: "Skull win",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeEscape}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[3], Card: Card{Type: CardTypeSuitGreen, Value: 14}},
				},
			},
			Want: players[2],
		},
		{
			Name: "Skull win black",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeEscape}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			Want: players[2],
		},
		{
			Name: "Skull win pirate",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			Want: players[2],
		},
		{
			Name: "pirate loose Skull",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypePirate}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			Want: players[0],
		},
		{
			Name: "Mermaid win Skull",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeMermaid}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			Want: players[2],
		},
		{
			Name: "Skull lose Mermaid",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypeSkullKing}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			Want: players[0],
		},
		{
			Name: "pirate win mermaid",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Player: players[2], Card: Card{Type: CardTypePirate}},
					{Player: players[3], Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			Want: players[2],
		},
		{
			Name: "only pirates",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypePirate}},
					{Player: players[1], Card: Card{Type: CardTypePirate}},
					{Player: players[2], Card: Card{Type: CardTypePirate}},
					{Player: players[3], Card: Card{Type: CardTypePirate}},
				},
			},
			Want: players[0],
		},
		{
			Name: "only mermaid",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypeMermaid}},
					{Player: players[2], Card: Card{Type: CardTypeSuitPurple, Value: 3}},
				},
			},
			Want: players[0],
		},
		{
			Name: "one each type",
			Trick: Trick{
				Table: []*Play{
					{Player: players[0], Card: Card{Type: CardTypeMermaid}},
					{Player: players[1], Card: Card{Type: CardTypePirate}},
					{Player: players[2], Card: Card{Type: CardTypeSkullKing}},
				},
			},
			Want: players[0],
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
					{Card: Card{Type: CardTypeSuitGreen}},
					{Card: Card{Type: CardTypeSuitGreen}},
					{Card: Card{Type: CardTypeSuitGreen}},
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
		{
			name: "Escape",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeEscape}},
					{Card: Card{Type: CardTypeSuitYellow}},
					{Card: Card{Type: CardTypeSuitBlack}},
				},
			},
			want: CardTypeSuitYellow,
		},
		{
			name: "Escape two times",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeEscape}},
					{Card: Card{Type: CardTypeEscape}},
					{Card: Card{Type: CardTypeSuitBlack}},
				},
			},
			want: CardTypeSuitBlack,
		},
		{
			name: "Escape three times",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeEscape}},
					{Card: Card{Type: CardTypeEscape}},
					{Card: Card{Type: CardTypeEscape}},
				},
			},
			want: CardTypeNone,
		},
		{
			name: "Escape pirate suit",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeEscape}},
					{Card: Card{Type: CardTypePirate}},
					{Card: Card{Type: CardTypeSuitGreen}},
				},
			},
			want: CardTypeNone,
		},
		{
			name: "Mermaid pirate suit",
			trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeMermaid}},
					{Card: Card{Type: CardTypePirate}},
					{Card: Card{Type: CardTypeSuitGreen}},
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
