### Least count

- Number of players(2)
- Each player will get 5 cards
- We should have 2-10 numbers and jack, queen, king, Ace
- We should have Spades, hearts, diamonds, clubs
- Shuffle the cards
- One random card as joker and running card
- we need discard pile
- We can pick the latest of discard pile/ from the deck
- Adding the score of one player
- Comparing score of all player to get to know who is least
- Different decks
- Joker is +1
- Match the same number from the discard pile
- More than 3 cards of anything is a match
- If someone calls least count 

## Pseudo classes

- Game
  - Should know about deck
  - Info about all the rounds
  - The rules and regulations of the game
  - It holds the players
  - It checks for the least count
  - It should hold the discard pile

- GameController
  - It is a bridge between Game and Player
  - It manages the turn of a player

- Player Object
  - Cards
  - Score for each round
  - Should allow to play a card
  - Should allow to pick a card
  - Has control to say least count

- Playing card distributor(Module)
  - Number of players, Number of cards per player, Deck
  - Should shuffle the cards
  - Should Distribute to the players
  - It knows what a deck is
  - It knows how to shuffle
  - Should return list of cards that each player holds

- Least count playing card Deck
  - How many jokers
  - There are 52 cards + jokers
  - It should know spades, clubs, diamonds, hearts
  - It should define all 13 cards
  - Should assign the values of each card
  - Should return the card deck

- Have your own immutuable class
  - Copies
  - References