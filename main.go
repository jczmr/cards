package main

func main() {

	cards := newDeck()

	hand, remaininCards := deal(cards, 5)

	hand.Println()
	remaininCards.Println()
}
