package main

import "testing"

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
