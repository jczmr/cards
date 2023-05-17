package main

func main() {

	/*
		cards := newDeck()

		//hand, remaininCards := deal(cards, 5)
		//hand.Println()
		//remaininCards.Println()

		fmt.Println(cards.toString())
		cards.saveToFile("my_cards")
	*/

	/*
		cards := newDeckFromFile("my_cards")
		cards.Println()
	*/

	cards := newDeck()
	cards.shuffle()
	cards.Println()

}
