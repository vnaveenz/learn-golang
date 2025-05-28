package main

func main()  {
	fn := "test.txt"
	cards := newDeck()
	cards.shuffle()
	cards.writeToFile(fn)
	cards.print()
	// dff := newDeckFromFile(fn)
	// dff.print()
	// deal, remainingDeck := deal(cards, 5)
	// deal.print()
	// remainingDeck.print()
}
