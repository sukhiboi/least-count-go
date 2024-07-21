package deck

func CreateCard(suit Suit, symbol Symbol, value int) Card {
	return Card{suit, symbol, value, false}
}

func Create() [54]Card {

	var cards [54]Card

	for suitIndex, suit := range suits {
		for symbolIndex, symbol := range symbols {
			card := CreateCard(suit, symbol, symbolValues[symbol])
			cards[(suitIndex*13)+symbolIndex] = card
		}
	}

	joker := Card{NONE, JOKER, 0, true}

	cards[52] = joker
	cards[53] = joker

	return cards
}
