package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(d))
	}

}

func TestDeal(t *testing.T) {
	cards := newDeck()
	hand, remaininCards := deal(cards, 5)

	if len(remaininCards) != 11 {
		hand.Println()
		remaininCards.Println()
		t.Errorf("Expected remaininCards lenght of 5, but got %v", len(remaininCards))
	}
}

func TestToString(t *testing.T) {
	cards := newDeck()

	if len(cards.toString()) == 0 {
		t.Errorf("Expected remaininCards lenght of 5, but got %v", len(cards))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
