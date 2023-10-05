package skullking

import (
	"reflect"
	"testing"
)

type TableTestWinner struct {
	Name  string
	Trick Trick

	WantPlayer string
	WantCard   Card
}

func TestTrick_Winner(t *testing.T) {
	table := []TableTestWinner{
		{
			Name: "higher card wins the trick",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{PlayerName: "2", Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			WantPlayer: "3",
		},
		{
			Name: "Different suit looses",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{PlayerName: "2", Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitGreen, Value: 14}},
				},
			},
			WantPlayer: "2",
		},
		{
			Name: "Triumph wins against suit",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{PlayerName: "2", Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitBlack, Value: 1}},
				},
			},
			WantPlayer: "3",
		},
		{
			Name: "Highest triumph wins",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeSuitBlack, Value: 1}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitBlack, Value: 12}},
					{PlayerName: "2", Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitBlack, Value: 10}},
				},
			},
			WantPlayer: "1",
		},
		{
			Name: "Pirate wins",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 12}},
					{PlayerName: "2", Card: Card{Type: CardTypePirate}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitBlack, Value: 10}},
				},
			},
			WantPlayer: "2",
		},
		{
			Name: "Second pirate looses",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{PlayerName: "1", Card: Card{Type: CardTypePirate}},
					{PlayerName: "2", Card: Card{Type: CardTypePirate}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitBlack, Value: 10}},
				},
			},
			WantPlayer: "1",
		},
		{
			Name: "Escape looses",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeEscape}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{PlayerName: "2", Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitGreen, Value: 14}},
				},
			},
			WantPlayer: "2",
		},
		{
			Name: "Escape everywhere",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeEscape}},
					{PlayerName: "1", Card: Card{Type: CardTypeEscape}},
					{PlayerName: "2", Card: Card{Type: CardTypeEscape}},
					{PlayerName: "3", Card: Card{Type: CardTypeEscape}},
				},
			},
			WantPlayer: "0",
		},
		{
			Name: "Escape looses black",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeEscape}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{PlayerName: "2", Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			WantPlayer: "3",
		},
		{
			Name: "Skull win",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeEscape}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{PlayerName: "2", Card: Card{Type: CardTypeSkullKing}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitGreen, Value: 14}},
				},
			},
			WantPlayer: "2",
		},
		{
			Name: "Skull win black",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeEscape}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{PlayerName: "2", Card: Card{Type: CardTypeSkullKing}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			WantPlayer: "2",
		},
		{
			Name: "Skull win pirate",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypePirate}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{PlayerName: "2", Card: Card{Type: CardTypeSkullKing}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			WantPlayer: "2",
		},
		{
			Name: "pirate loose Skull",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeSkullKing}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{PlayerName: "2", Card: Card{Type: CardTypePirate}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			WantPlayer: "0",
		},
		{
			Name: "Mermaid win Skull",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeSkullKing}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{PlayerName: "2", Card: Card{Type: CardTypeMermaid}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			WantPlayer: "2",
		},
		{
			Name: "Skull lose Mermaid",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeMermaid}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{PlayerName: "2", Card: Card{Type: CardTypeSkullKing}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			WantPlayer: "0",
		},
		{
			Name: "pirate win mermaid",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeMermaid}},
					{PlayerName: "1", Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{PlayerName: "2", Card: Card{Type: CardTypePirate}},
					{PlayerName: "3", Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			WantPlayer: "2",
		},
		{
			Name: "only pirates",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypePirate}},
					{PlayerName: "1", Card: Card{Type: CardTypePirate}},
					{PlayerName: "2", Card: Card{Type: CardTypePirate}},
					{PlayerName: "3", Card: Card{Type: CardTypePirate}},
				},
			},
			WantPlayer: "0",
		},
		{
			Name: "only mermaid",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeMermaid}},
					{PlayerName: "1", Card: Card{Type: CardTypeMermaid}},
					{PlayerName: "2", Card: Card{Type: CardTypeSuitPurple, Value: 3}},
				},
			},
			WantPlayer: "0",
		},
		{
			Name: "one each type",
			Trick: Trick{
				Table: []*Play{
					{PlayerName: "0", Card: Card{Type: CardTypeMermaid}},
					{PlayerName: "1", Card: Card{Type: CardTypePirate}},
					{PlayerName: "2", Card: Card{Type: CardTypeSkullKing}},
				},
			},
			WantPlayer: "0",
		},
	}
	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			got := tt.Trick.Winner()
			if !reflect.DeepEqual(got.PlayerName, tt.WantPlayer) {
				t.Errorf("Trick.Winner() = %#v, want %#v", got.PlayerName, tt.WantPlayer)
			}
		})
	}

}

type TableTestLeading struct {
	name  string
	trick Trick
	want  CardType
}

func TestTrick_Leading(t *testing.T) {
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

type TableTestBonus struct {
	Name  string
	Trick Trick

	Want int
}

func TestTrick_Bonus_Points(t *testing.T) {

	table := []TableTestBonus{
		{
			Name: "low numbers",
			Trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeSuitYellow, Value: 1}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 2}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 3}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: 0,
		},
		{
			Name: "max numbers",
			Trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeSuitYellow, Value: 14}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{Card: Card{Type: CardTypeSuitGreen, Value: 14}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 4}},
				},
			},
			Want: 20,
		},
		{
			Name: "with low Black",
			Trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeSuitYellow, Value: 14}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{Card: Card{Type: CardTypeSuitGreen, Value: 14}},
					{Card: Card{Type: CardTypeSuitBlack, Value: 4}},
				},
			},
			Want: 20,
		},
		{
			Name: "with max Black",
			Trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeSuitYellow, Value: 14}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{Card: Card{Type: CardTypeSuitGreen, Value: 14}},
					{Card: Card{Type: CardTypeSuitBlack, Value: 14}},
				},
			},
			Want: 40,
		},
		{
			Name: "only one pirate",
			Trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeSuitYellow, Value: 14}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{Card: Card{Type: CardTypeSuitGreen, Value: 14}},
					{Card: Card{Type: CardTypeSuitBlack, Value: 14}},
					{Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 40,
		},
		{
			Name: "one pirate and one mermaid",
			Trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeMermaid}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{Card: Card{Type: CardTypeSuitGreen, Value: 14}},
					{Card: Card{Type: CardTypeSuitBlack, Value: 14}},
					{Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 50,
		},
		{
			Name: "one pirate and two mermaids",
			Trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeMermaid}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{Card: Card{Type: CardTypeSuitGreen, Value: 14}},
					{Card: Card{Type: CardTypeMermaid}},
					{Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 50,
		},
		{
			Name: "one pirate and skullking",
			Trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeSkullKing}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{Card: Card{Type: CardTypeSuitGreen, Value: 14}},
					{Card: Card{Type: CardTypeSuitBlack, Value: 14}},
					{Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 60,
		},
		{
			Name: "two pirates and skullking",
			Trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeSkullKing}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{Card: Card{Type: CardTypeSuitGreen, Value: 14}},
					{Card: Card{Type: CardTypePirate}},
					{Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 70,
		},
		{
			Name: "one mermaid and skullking",
			Trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeSkullKing}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{Card: Card{Type: CardTypeSuitGreen, Value: 14}},
					{Card: Card{Type: CardTypeSuitBlack, Value: 14}},
					{Card: Card{Type: CardTypeMermaid}},
				},
			},
			Want: 70,
		},
		{
			Name: "one mermaid, one pirate and skullking",
			Trick: Trick{
				Table: []*Play{
					{Card: Card{Type: CardTypeSkullKing}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 13}},
					{Card: Card{Type: CardTypeSuitYellow, Value: 14}},
					{Card: Card{Type: CardTypeMermaid}},
					{Card: Card{Type: CardTypePirate}},
				},
			},
			Want: 50,
		},
	}

	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			got := tt.Trick.Bonus()
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("Trick.Bonus() = %v, want %v", got, tt.Want)
			}
		})
	}
}
