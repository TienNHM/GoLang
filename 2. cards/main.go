package main

import "fmt"

var deckSize int = 10

func main() {
	// var card string = "Ace of Spades"
	card := newCard()

	fmt.Println(card)
	fmt.Println(deckSize)
}

func newCard() string {
	return "Ace of Spades"
}
