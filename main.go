package main

import "fmt"

func main() {

	cards := newDeck()

	//hand, remaininCards := deal(cards, 5)
	//hand.Println()
	//remaininCards.Println()

	fmt.Println(cards.toString())
}
