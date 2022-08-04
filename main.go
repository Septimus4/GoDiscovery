package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	hand, remaining := cards.deal(5)
	fmt.Println("\nRemaining")
	remaining.print()
	fmt.Println("\nHand")
	hand.print()

	remaining.shuffle()
	fmt.Println("\nShuffled")
	remaining.print()
}
