package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// fmt.Println()
	// remainingDeck.print()
	// fmt.Println()
	// cards.print()

	// cards := newDeck()
	//fmt.Println(cards.toString())

	// cards.saveToFile("myCards.doc")
	// cards.saveToFile("C:\\Users\\user\\Desktop\\ua.txt")
	// cards.saveToFile("C:\\Users\\user\\Desktop\\Go\\helloWorld\\uaCards")

	cards := readFromFile("myCards.doc")
	cards.print()

	cards.shuffle()
	cards.print()
}
