package deck

func Create() [54]Card {

	var cards [54]Card

	for suitIndex, suit := range suits {
		for symbolIndex, symbol := range symbols {
			card := Card{suit, symbol, symbolValues[symbol]}
			cards[(suitIndex*13)+symbolIndex] = card
		}
	}

	joker := Card{NONE, JOKER, 0}

	cards[52] = joker
	cards[53] = joker

	return cards
}
