package main

func main() {
	cards := newDeck()
	cards.print()

	hand, remaining := cards.deal(5)
	remaining.print()
	hand.print()
}
