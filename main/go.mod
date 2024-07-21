module main

go 1.22.5

replace least_count/deck => ../deck

require least_count/setup v0.0.0-00010101000000-000000000000

require (
	least_count/deck v0.0.0-00010101000000-000000000000 // indirect
	least_count/game v0.0.0-00010101000000-000000000000 // indirect
	least_count/player v0.0.0-00010101000000-000000000000 // indirect
)

replace least_count/setup => ../setup

replace least_count/player => ../player

replace least_count/game => ../game
