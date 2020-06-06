package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52 cards (found %v).", len(d))
	}
	if d[0] != "♣A" {
		t.Errorf("Expected first card to be ♣A (found %v).", d[0])
	}
	if d[len(d)-1] != "♠K" {
		t.Errorf("Expected last card to be ♠K (found %v).", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove(".decktest")
	d := newDeck()
	d.saveToFile(".decktest")
	loadedDeck := newDeckFromFile(".decktest")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards (found %v).", len(loadedDeck))
	}
	os.Remove(".decktest")
}
