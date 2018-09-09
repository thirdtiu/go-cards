package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	// two variables since deal function returns 2 value of type deck, first will be assigned to hand and 2nd will be assigned to remainingCards

	hand.print()
	remainingCards.print()
}
