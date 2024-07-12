package main

import (
	"fmt"
	"least_count/deck"
	"least_count/playingCardDistributor"
)

func main() {
	cards := deck.Create()
	fmt.Println(playingCardDistributor.Shuffle(cards))
}
