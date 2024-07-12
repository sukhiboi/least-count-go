package playingCardDistributor

import (
	"least_count/deck"
	"least_count/player"
	"math/rand"
)

func Shuffle(cards [54]deck.Card) [54]deck.Card {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func Distribute(players []player.Player, cards [54]deck.Card, noOfCards int) ([]player.Player, []deck.Card) {
	for index := 0; index < noOfCards; index++ {
		for playerCount := 0; playerCount < len(players); playerCount++ {
			players[playerCount].Cards = append(players[playerCount].Cards, cards[index*noOfCards])
			// remove the distributed cards from cards object and return remaining cards
		}
	}

	return players, cards[:]
}
