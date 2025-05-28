package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"heart", "spade", "diamond", "claire"}
	cardValues := []string{"Ace", "One", "two", "three", "four"}
	for _,cs := range cardSuites {
		for _, cv := range cardValues {
			cards = append(cards, cv+" of "+cs)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	dealCards := d[:handSize]
	remainingCards := d[handSize:]
	return dealCards, remainingCards
}

func (d deck) writeToFile (fn string) {
	sd := []string(d)
	bs := []byte(strings.Join(sd, ","))
	os.WriteFile(fn, bs, 0666)
}

func newDeckFromFile (fn string) (d deck) {
	bs, err := os.ReadFile(fn)
	str := strings.Split(string(bs), ",")
	fd := deck(str)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return fd
}

func (d deck) shuffle() {
	for i := range d {
		newPos := rand.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
