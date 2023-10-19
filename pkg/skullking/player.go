package skullking

type Player struct {
	Name        string
	Hand        Cards
	Scoresheets []Scoresheet
}



/*func test_player() {
	p := Player{
		Name: "Sirena",
		Hand: Cards{Card{Type: CardTypeEscape}, Card{Value: 7, Type: CardTypeSuitParrot}},
		Scoresheets: []Scoresheet{
			Scoresheet{
				Round:       1,
				CardsDealt:  1, ///todo sustituible por una función
				Bid:         1,
				Result:      1,
				BidPoints:   20, //todo
				Bonus:       0,
				RoundPoints: 20, //todo hacer función para calcular estos puntos
				RunnigTotal: 20, //todo
			},
			Scoresheet{
				Round:       2,
				CardsDealt:  2,
				Bid:         0,
				Result:      0,
				BidPoints:   20,
				Bonus:       0,
				RoundPoints: 20,
				RunnigTotal: 40,
			},
		},
	}

}*/
