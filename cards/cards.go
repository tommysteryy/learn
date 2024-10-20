package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	card, suit string
}

func (c Card) read() {
	fmt.Println(c.card + " of " + c.suit)
}

const NUM_CARDS = 13
const NUM_SUITS = 4

// All card names
var cardNames = [NUM_CARDS]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
var suitNames = [NUM_SUITS]string{"Spades", "Hearts", "Diamonds", "Clubs"}

func randomCard() Card {
	card := rand.Intn(NUM_CARDS)
	suit := rand.Intn(NUM_SUITS)

	return Card{cardNames[card], suitNames[suit]}
}
