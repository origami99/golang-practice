package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	deckLenExpected := 52
	deckLenActual := len(d)
	if deckLenExpected != deckLenActual {
		t.Errorf("Expected deck length of %v but got %v", deckLenExpected, deckLenActual)
	}

	firstCardExpected := "Ace of Spades"
	firstCardActual := d[0]
	if firstCardExpected != firstCardActual {
		t.Errorf("Expected first card of %v but got %v", firstCardExpected, firstCardActual)
	}

	lastCardExpected := "King of Clubs"
	lastCardActual := d[len(d)-1]
	if lastCardExpected != lastCardActual {
		t.Errorf("Expected first card of %v but got %v", lastCardExpected, lastCardActual)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
