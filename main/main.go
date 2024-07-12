package main

import (
	"least_count/deck"
	"least_count/player"
	"least_count/playingCardDistributor"
)

func main() {
	cards := deck.Create()
	firstPlayer := player.Player{Name: "John", Cards: []deck.Card{}, Score: 0}
	secondPlayer := player.Player{Name: "Robot", Cards: []deck.Card{}, Score: 0}
	players := []player.Player{firstPlayer, secondPlayer}
	shuffledCards := playingCardDistributor.Shuffle(cards)
	playingCardDistributor.Distribute(players, shuffledCards, 5)
}
