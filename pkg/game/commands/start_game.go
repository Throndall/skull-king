package commands

import (
	"github.com/metalblueberry/skull-king/pkg/game"
	"github.com/metalblueberry/skull-king/pkg/skullking"
)

type StartGame struct {
	Seed int64
}

func (c StartGame) Execute(state *game.State) error {
	state.Deck = skullking.NewDeck(c.Seed)
	state.Deck.Put(skullking.Card{Number: 1, Type: skullking.CardTypeSuitParrot, Value: 1},
		skullking.Card{Number: 2, Type: skullking.CardTypeSuitParrot, Value: 2},
		skullking.Card{Number: 3, Type: skullking.CardTypeSuitParrot, Value: 3},
		skullking.Card{Number: 4, Type: skullking.CardTypeSuitParrot, Value: 4},
		skullking.Card{Number: 5, Type: skullking.CardTypeSuitParrot, Value: 5},
		skullking.Card{Number: 6, Type: skullking.CardTypeSuitParrot, Value: 6},
		skullking.Card{Number: 7, Type: skullking.CardTypeSuitParrot, Value: 7},
		skullking.Card{Number: 8, Type: skullking.CardTypeSuitParrot, Value: 8},
		skullking.Card{Number: 9, Type: skullking.CardTypeSuitParrot, Value: 9},
		skullking.Card{Number: 10, Type: skullking.CardTypeSuitParrot, Value: 10},
		skullking.Card{Number: 11, Type: skullking.CardTypeSuitParrot, Value: 11},
		skullking.Card{Number: 12, Type: skullking.CardTypeSuitParrot, Value: 12},
		skullking.Card{Number: 13, Type: skullking.CardTypeSuitParrot, Value: 13},
		skullking.Card{Number: 14, Type: skullking.CardTypeSuitParrot, Value: 14},
		skullking.Card{Number: 15, Type: skullking.CardTypeSuitTreasureChest, Value: 1},
		skullking.Card{Number: 16, Type: skullking.CardTypeSuitTreasureChest, Value: 2},
		skullking.Card{Number: 17, Type: skullking.CardTypeSuitTreasureChest, Value: 3},
		skullking.Card{Number: 18, Type: skullking.CardTypeSuitTreasureChest, Value: 4},
		skullking.Card{Number: 19, Type: skullking.CardTypeSuitTreasureChest, Value: 5},
		skullking.Card{Number: 20, Type: skullking.CardTypeSuitTreasureChest, Value: 6},
		skullking.Card{Number: 21, Type: skullking.CardTypeSuitTreasureChest, Value: 7},
		skullking.Card{Number: 22, Type: skullking.CardTypeSuitTreasureChest, Value: 8},
		skullking.Card{Number: 23, Type: skullking.CardTypeSuitTreasureChest, Value: 9},
		skullking.Card{Number: 24, Type: skullking.CardTypeSuitTreasureChest, Value: 10},
		skullking.Card{Number: 25, Type: skullking.CardTypeSuitTreasureChest, Value: 11},
		skullking.Card{Number: 26, Type: skullking.CardTypeSuitTreasureChest, Value: 12},
		skullking.Card{Number: 27, Type: skullking.CardTypeSuitTreasureChest, Value: 13},
		skullking.Card{Number: 28, Type: skullking.CardTypeSuitTreasureChest, Value: 14},
		skullking.Card{Number: 29, Type: skullking.CardTypeSuitJollyRoger, Value: 1},
		skullking.Card{Number: 30, Type: skullking.CardTypeSuitJollyRoger, Value: 2},
		skullking.Card{Number: 31, Type: skullking.CardTypeSuitJollyRoger, Value: 3},
		skullking.Card{Number: 32, Type: skullking.CardTypeSuitJollyRoger, Value: 4},
		skullking.Card{Number: 33, Type: skullking.CardTypeSuitJollyRoger, Value: 5},
		skullking.Card{Number: 34, Type: skullking.CardTypeSuitJollyRoger, Value: 6},
		skullking.Card{Number: 35, Type: skullking.CardTypeSuitJollyRoger, Value: 7},
		skullking.Card{Number: 36, Type: skullking.CardTypeSuitJollyRoger, Value: 8},
		skullking.Card{Number: 37, Type: skullking.CardTypeSuitJollyRoger, Value: 9},
		skullking.Card{Number: 38, Type: skullking.CardTypeSuitJollyRoger, Value: 10},
		skullking.Card{Number: 39, Type: skullking.CardTypeSuitJollyRoger, Value: 11},
		skullking.Card{Number: 40, Type: skullking.CardTypeSuitJollyRoger, Value: 12},
		skullking.Card{Number: 41, Type: skullking.CardTypeSuitJollyRoger, Value: 13},
		skullking.Card{Number: 42, Type: skullking.CardTypeSuitJollyRoger, Value: 14},
		skullking.Card{Number: 43, Type: skullking.CardTypeSuitPirateMap, Value: 1},
		skullking.Card{Number: 44, Type: skullking.CardTypeSuitPirateMap, Value: 2},
		skullking.Card{Number: 45, Type: skullking.CardTypeSuitPirateMap, Value: 3},
		skullking.Card{Number: 46, Type: skullking.CardTypeSuitPirateMap, Value: 4},
		skullking.Card{Number: 47, Type: skullking.CardTypeSuitPirateMap, Value: 5},
		skullking.Card{Number: 48, Type: skullking.CardTypeSuitPirateMap, Value: 6},
		skullking.Card{Number: 49, Type: skullking.CardTypeSuitPirateMap, Value: 7},
		skullking.Card{Number: 50, Type: skullking.CardTypeSuitPirateMap, Value: 8},
		skullking.Card{Number: 51, Type: skullking.CardTypeSuitPirateMap, Value: 9},
		skullking.Card{Number: 52, Type: skullking.CardTypeSuitPirateMap, Value: 10},
		skullking.Card{Number: 53, Type: skullking.CardTypeSuitPirateMap, Value: 11},
		skullking.Card{Number: 54, Type: skullking.CardTypeSuitPirateMap, Value: 12},
		skullking.Card{Number: 55, Type: skullking.CardTypeSuitPirateMap, Value: 13},
		skullking.Card{Number: 56, Type: skullking.CardTypeSuitPirateMap, Value: 14},

		skullking.Card{Number: 57, Type: skullking.CardTypePirate},
		skullking.Card{Number: 58, Type: skullking.CardTypePirate},
		skullking.Card{Number: 59, Type: skullking.CardTypePirate},
		skullking.Card{Number: 60, Type: skullking.CardTypePirate},
		skullking.Card{Number: 61, Type: skullking.CardTypePirate},
		skullking.Card{Number: 62, Type: skullking.CardTypeTigress},
		skullking.Card{Number: 63, Type: skullking.CardTypeSkullKing},
		skullking.Card{Number: 64, Type: skullking.CardTypeMermaid},
		skullking.Card{Number: 65, Type: skullking.CardTypeMermaid},
		skullking.Card{Number: 66, Type: skullking.CardTypeEscape},
		skullking.Card{Number: 67, Type: skullking.CardTypeEscape},
		skullking.Card{Number: 68, Type: skullking.CardTypeEscape},
		skullking.Card{Number: 69, Type: skullking.CardTypeEscape},
		skullking.Card{Number: 70, Type: skullking.CardTypeEscape})
	state.Deck.Shufle()

	for i := 0; i < len(state.Players); i++ {
		state.Players[i].Hand = state.Deck.Draw(1)
	}

	return nil
}
