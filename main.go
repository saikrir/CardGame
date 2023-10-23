package main

import "fmt"

func main() {
	fmt.Println("Welcome to Card Game")
	deck := newDeck()
	deck.shuffle()
	//deck.print()
	//dealt, _ := deal(deck, 40)
	//dealt.print()
	err := deck.saveToFile("Deck.txt")
	if err != nil {
		fmt.Println("File Could not be Written ", err)
	} else {
		fmt.Println("Written to a file")
	}

	serialzedDeck := newDeckFromFile("Deck.txt")
	serialzedDeck.print()
}
