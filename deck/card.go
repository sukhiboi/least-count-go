package deck

type Suit string
type Symbol string

const (
	SPADE   Suit = "spade"
	HEART   Suit = "heart"
	CLUB    Suit = "club"
	DIAMOND Suit = "diamond"
	NONE    Suit = "none"
)

const (
	KING  Symbol = "king"
	QUEEN Symbol = "queen"
	JACK  Symbol = "jack"
	TEN   Symbol = "ten"
	NINE  Symbol = "nine"
	EIGHT Symbol = "eight"
	SEVEN Symbol = "seven"
	SIX   Symbol = "six"
	FIVE  Symbol = "five"
	FOUR  Symbol = "four"
	THREE Symbol = "three"
	TWO   Symbol = "two"
	ACE   Symbol = "ace"
	JOKER Symbol = "joker"
)

type Card struct {
	suit    Suit
	Symbol  Symbol
	value   int
	IsJoker bool
}

var symbols = []Symbol{
	KING,
	QUEEN,
	JACK,
	TEN,
	NINE,
	EIGHT,
	SEVEN,
	SIX,
	FIVE,
	FOUR,
	THREE,
	TWO,
	ACE,
}

var suits = []Suit{
	SPADE,
	HEART,
	CLUB,
	DIAMOND,
}

var symbolValues = map[Symbol]int{
	KING:  10,
	QUEEN: 10,
	JACK:  10,
	TEN:   10,
	NINE:  9,
	EIGHT: 8,
	SEVEN: 7,
	SIX:   6,
	FIVE:  5,
	FOUR:  4,
	THREE: 3,
	TWO:   2,
	ACE:   1,
	JOKER: 0,
}
