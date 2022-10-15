package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades"
	card := newCard()

	cards := []string{"A", "B"}

	numbers := []int{1, 2}

	cards = append(cards, "C", "D")
	x := len(cards)

	fmt.Println(card)
	fmt.Println(cards)
	fmt.Println(x)
	fmt.Println(numbers)
}

func newCard() string {
	return "Ace of Spades"
}
