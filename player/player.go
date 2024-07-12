package player

import "least_count/deck"

type Player struct {
	Name  string      `default:""`
	Cards []deck.Card `default:"[]"`
	Score uint        `default:"0"`
}
