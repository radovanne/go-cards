package main

func main() {

	cardDeck := newDeck()

	hand, remainingDeck := deal(cardDeck, 5)

	hand.print()
	remainingDeck.print()

}
