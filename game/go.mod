module least_count/game

go 1.22.5

replace least_count/player => ../player

replace least_count/deck => ../deck

require least_count/player v0.0.0-00010101000000-000000000000

require least_count/deck v0.0.0-00010101000000-000000000000 // indirect
