package main

import (
	"fmt"
	"least_count/setup"
)

func main() {
	game := setup.Setup(2, 3)
	fmt.Println((game.Players))
	fmt.Println((game.Joker))
}
