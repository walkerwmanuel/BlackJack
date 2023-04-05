package logic

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"walkerwmanuel/blackjack/types"
)

var (
	suites = []string{"Hearts", "Clubs", "Diamonds", "Spades"}
	values = []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}
)

// NewDeck - creates a deck with 52 cards based off map of types.Card
func NewDeck() []types.Card {
	//Makes an empty map of 52 type card
	d := make([]types.Card, 52)

	index := 0
	//Itteration through all 52 types of cards and stores it
	for i := range suites {
		for j := range values {
			d[index] = newCard(suites[i], values[j])
			index++
		}
	}
	return d
}

// Shuffle - Shuffles through deck and places the cards in a random order
func ShuffleDeck(d []types.Card) {
	if len(d) == 0 {
		return
	}

	// iterate through each card in Cards of deck
	for i := range d {
		// pull out card of current iteration index 0...len(d.Cards)
		card := d[i]
		// creates a new (random) position from 0...len(d.Cards)
		newPos, _ := rand.Int(rand.Reader, big.NewInt(int64(len(d))))
		// convert newPos to int
		newPosInt := newPos.Uint64()
		// pull out card in new position
		otherCard := d[newPosInt]
		// swap them
		d[i] = otherCard
		d[newPosInt] = card
	}
}

// PrintDeck prints the contents of a deck
func PrintDeck(d []types.Card) {
	fmt.Println()
	for i := range d {
		fmt.Printf("[%s of %s] ", d[i].Suite, d[i].Value)
	}
	fmt.Println()
}

// NewCard - makes a new card
func newCard(v, s string) types.Card {
	c := types.Card{
		Suite:  s,
		Filler: "of",
		Value:  v,
	}
	return c
}

func DealCard(d []types.Card) ([]types.Card, []types.Card) {
	hand := d[0:1]
	deck := d[1:]
	return deck, hand
}
