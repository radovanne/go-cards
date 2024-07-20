package main

import (
	"fmt"
	"strings"
)

type deck []string

/*This function prints every card from the deck it was called upon.*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
This function deals cards from the deck.
It returns 2 values 'hand' and 'remaining set of cards' both of the type 'deck'
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*
Helper function that converts whole deck slice to
a string separated by commas ","
*/
func (d deck) toString() string {
	return strings.Join(d, ",")
}

/*This function creates a new deck of cards..*/
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Twelve", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
