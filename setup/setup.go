package setup

import (
	"least_count/deck"
	"least_count/game"
	"least_count/player"
	"math/rand"
)

func shuffle(cards [54]deck.Card) [54]deck.Card {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func distribute(numberOfPlayers int, cards [54]deck.Card, numberOfCardsPerPlayer int) ([][]deck.Card, []deck.Card) {

	distributedCards := make([][]deck.Card, numberOfPlayers)
	runningCardIndex := len(cards)

	for round := 0; round < numberOfCardsPerPlayer; round++ {
		for player := 0; player < numberOfPlayers; player++ {

			runningCardIndex--

			if distributedCards[player] == nil {
				distributedCards[player] = []deck.Card{}
			}

			distributedCards[player] = append(distributedCards[player], cards[runningCardIndex])
		}
	}

	return distributedCards, cards[0:runningCardIndex]
}

func updateJoker(cards []deck.Card, joker deck.Symbol) []deck.Card {
	for index, card := range cards {
		if card.Symbol == joker {
			cards[index].IsJoker = true
		}
	}
	return cards
}

func Setup(numberOfPlayers int, numberOfCardsPerPlayer int) game.Game {
	players := make([]player.Player, numberOfPlayers)
	cards := deck.Create()
	shuffledCards := shuffle(cards)
	distributedCards, remainingCards := distribute(numberOfPlayers, shuffledCards, numberOfCardsPerPlayer)

	discardedPile := []deck.Card{remainingCards[0]}
	remainingCards = remainingCards[1:]
	randomNumber := rand.Intn(len(remainingCards) - 1)
	joker := remainingCards[randomNumber]
	remainingCards = updateJoker(remainingCards, joker.Symbol)

	for index, oneSet := range distributedCards {
		players[index].Cards = updateJoker(oneSet, joker.Symbol)
	}

	remainingCards = append(remainingCards[:randomNumber], remainingCards[randomNumber+1:]...)
	return game.Game{Players: players, RemainingCards: remainingCards, DiscardedPile: discardedPile, Joker: joker.Symbol}
}
