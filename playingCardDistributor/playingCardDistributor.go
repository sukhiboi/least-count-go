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

func Distribute(numberOfPlayers int, cards [54]deck.Card, numberOfCardsPerPlayer int) ([][]deck.Card, []deck.Card) {

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

