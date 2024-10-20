package main

import (
	"fmt"
	"math/rand"
)

type Deck struct {
	cards []Card // How do I make this fixed size?
}

func (d Deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d Deck) read() {
	fmt.Println("There are", len(d.cards), "cards in the deck.")
	for _, card := range d.cards {
		card.read()
	}
}

// dealHand deals a hand of cards from the deck.
// It takes the size of the hand as a parameter and returns a slice of Card.
// The size parameter specifies the number of cards to be dealt.
// IMPORTANT: The function modifies the deck by removing the dealt cards.
func (d *Deck) dealHand(size int) []Card {

	if size > len(d.cards) {
		size = len(d.cards)
		fmt.Println("Not enough cards in the deck. Dealing", size, "cards instead.")
	}

	hand := make([]Card, size)
	for i := 0; i < size; i++ {
		hand[i] = d.cards[i]
	}
	d.cards = d.cards[size:]
	return hand
}

func newDeck() Deck {
	deck := make([]Card, NUM_CARDS*NUM_SUITS)
	for i := 0; i < NUM_CARDS; i++ {
		for j := 0; j < NUM_SUITS; j++ {
			deck[i*NUM_SUITS+j] = Card{cardNames[i], suitNames[j]}
		}
	}

	return Deck{cards: deck}
}
