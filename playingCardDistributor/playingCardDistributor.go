package playingCardDistributor

import (
	"least_count/deck"
	"math/rand"
)

func Shuffle(cards [54]deck.Card) [54]deck.Card {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}
