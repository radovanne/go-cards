package main

import (
	"fmt"
	"os"
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

/*
This function saves deck of cards to a file.
*/
func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

/*
This function reads deck of cards from a file.
*/
func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ",")
	return deck(ss)
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
