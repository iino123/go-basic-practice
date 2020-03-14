package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card is Ace of Spades, but got %s", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDockFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
