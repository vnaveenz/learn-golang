package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected 20, but got %v", len(d))
	}

	if d[0] != "Ace of heart" {
		t.Errorf("Expected ace of heart but got %v", d[0])
	}

	if d[len(d) - 1] != "four of claire" {
		t.Errorf("Expected four of claire but got %v", d[0])
	}
}

func TestWriteToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktestingfile")
	td := newDeck()
	td.writeToFile("_decktestingfile")
	ldd := newDeckFromFile("_decktestingfile")
	if len(ldd) != 20 {
		t.Errorf("Expected 20, but got %v", len(ldd))
	}
	os.Remove("_decktestingfile")
}
