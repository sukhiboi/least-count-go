package playingCardDistributor

import (
  "least_count/deck"
  "testing"
)

func TestHelloName(t *testing.T) {
  playerCount := 4
  numberOfCardsPerPlayer := 7
  givenDeck := deck.Create()

  var expected [][]deck.Card = make([][]deck.Card, playerCount)
  expected[0] = make([]deck.Card, numberOfCardsPerPlayer)
  expected[0][0] = deck.CreateCard(deck.NONE, deck.JOKER, 0)
  expected[0][1] = deck.CreateCard(deck.DIAMOND, deck.ACE, 1)
  expected[1] = make([]deck.Card, numberOfCardsPerPlayer)
  expected[1][0] = deck.CreateCard(deck.NONE, deck.JOKER, 0)
  expected[1][1] = deck.CreateCard(deck.SPADE, deck.TWO, 2)

  distributedCards, discardPile := Distribute(playerCount, givenDeck, numberOfCardsPerPlayer)
}
