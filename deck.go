package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

func newDeck() Deck {

	suites := []string{"♠", "♣", "♥", "◆"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	deck := Deck{}

	for _, suite := range suites {
		for _, cardValue := range cardValues {
			aCard := suite + " " + cardValue
			deck = append(deck, aCard)
		}
	}

	return deck
}

func deal(deck Deck, numCards int) (Deck, Deck) {
	return deck[:numCards], deck[numCards:]
}

func (d Deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d Deck) shuffle() {
	seedValue := time.Now().UnixMicro()
	source := rand.NewSource(seedValue)
	rand := rand.New(source)

	for i := range d {
		randomIndex := rand.Intn(len(d) - 1)
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
}

func (d Deck) toString() string {
	return strings.Join(d, ",")
}

func (d Deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) Deck {
	rawBytes, err := os.ReadFile(fileName)
	if err != nil {
		os.Exit(1)
	}

	serialzedDeck := string(rawBytes)
	return Deck(strings.Split(serialzedDeck, ","))
}
