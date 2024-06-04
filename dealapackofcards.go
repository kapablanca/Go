package piscine

import "fmt"

type player struct {
	name  string
	cards []int
}

// Create a player
func newPlayer(name string) *player {
	p := player{name: name}

	return &p
}

// Deal a player cards and remove them from deck
func dealDeck(p *player, deck *[]int, cards int) {
	p.cards = (*deck)[:cards]
	*deck = (*deck)[cards:]
}

// Function that takes  a pack of 12 cards as slice of integers
// and deals them evenly to 4 players. Each player must be printed
//
//	in a different line
func DealAPackOfCards(deck []int) {
	numberCards := len(deck)
	numberPlayers := 4
	cardsPlayer := numberCards / numberPlayers

	p1 := newPlayer("Player 1")
	p2 := newPlayer("Player 2")
	p3 := newPlayer("Player 3")
	p4 := newPlayer("Player 4")

	totalPlayers := []*player{p1, p2, p3, p4}
	// Deal cards to the players
	for _, player := range totalPlayers {
		dealDeck(player, &deck, cardsPlayer)
		// Print player name
		fmt.Printf("%s: ", player.name)
		// Print player hand
		for index, card := range player.cards {
			if index != 0 {
				fmt.Print(", ")
			}
			fmt.Print(card)
		}
		fmt.Print("\n")
	}
}
