package player

import "least_count/deck"

type Player struct {
	Cards []deck.Card `default:"[]"`
}
