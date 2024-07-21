package game

import (
	"least_count/deck"
	"least_count/player"
)

type Game struct {
	Players        []player.Player
	RemainingCards []deck.Card
	DiscardedPile  []deck.Card
	Joker          deck.Symbol
}
