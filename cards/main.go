package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the casino!")

	deck := newDeck()

	deck.read()

	fmt.Println("Shuffling...")
	// Shuffle
	deck.shuffle()

	deck.dealHand(50)
	secondHand := deck.dealHand(10)
	fmt.Println("Second hand:")
	for _, card := range secondHand {
		card.read()
	}

	deck.read()
}
