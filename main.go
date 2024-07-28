package main

func main() {
	cardDeck := newDeckFromFile("my_cards")
	cardDeck.shuffle()
	cardDeck.print()
}
