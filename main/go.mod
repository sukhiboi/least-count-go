module main

go 1.22.5

replace least_count/deck => ../deck

require (
	least_count/deck v0.0.0-00010101000000-000000000000
	least_count/playingCardDistributor v0.0.0-00010101000000-000000000000
	least_count/player v0.0.0-00010101000000-000000000000
)

replace least_count/playingCardDistributor => ../playingCardDistributor

replace least_count/player => ../player
