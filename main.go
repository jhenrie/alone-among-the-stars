package main

import (
	"fmt"
	"github.com/jhenrie/alone-among-the-stars/cards"
	"github.com/jhenrie/alone-among-the-stars/dice"
)

func main() {
	fmt.Printf("Hello World.\n")
	d6 := dice.NewDice(6)
	fmt.Printf("D6 Roll: %d \n", d6.Roll())
	fmt.Printf("D6 Roll: %d \n", d6.Roll())
	fmt.Printf("D6 Roll: %d \n", d6.Roll())

	fmt.Printf("D6 Roll List: %v \n", d6.RollN(10))

	d20 := dice.NewDice(20)
	fmt.Printf("D20 Roll: %d \n", d20.Roll())
	fmt.Printf("D20 Roll: %d \n", d20.Roll())
	fmt.Printf("D20 Roll: %d \n", d20.Roll())

	fmt.Printf("D20 Roll List: %v \n", d20.RollN(10))

	deck := cards.InitDeck()
	fmt.Printf("%s\n", deck)
	fmt.Printf("Deck Length - %d\n", len(deck))

	deck.Shuffle()
	fmt.Printf("%s\n", deck)

	card := deck.DrawOne()
	fmt.Printf("Card %s\n", card)
	fmt.Printf("Deck Length - %d\n", len(deck))

	deck.AddCard(card)
	fmt.Printf("Deck Length - %d\n", len(deck))
	fmt.Printf("%s\n", deck)

	card = deck.DrawOne()
	fmt.Printf("Card %s\n", card)
	fmt.Printf("Deck Length - %d\n", len(deck))


}
