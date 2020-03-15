package main

func main() {
	// cards := newDockFromFile("my_cardsaaa")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
