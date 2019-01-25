package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedDeckLen := 52
	actualDeckLen := len(d)
	if actualDeckLen != expectedDeckLen {
		t.Errorf("Expected deck length of ", expectedDeckLen, " but got ", actualDeckLen)
	}
}
