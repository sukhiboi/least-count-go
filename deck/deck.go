package deck

func Create() [54]Card {

	var cards [54]Card

	for i, suit := range suits {
		for j, symbol := range symbols {
			card := Card{suit, symbol, symbolValues[symbol]}
			cards[(i*13)+j] = card
		}
	}

	joker := Card{NONE, JOKER, 0}

	cards[52] = joker
	cards[53] = joker

	return cards
}
