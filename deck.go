package main

import (
	"fmt"
	"strings"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// create a new deck
func newDeck() deck {
	// empty deck
	cards := deck{}

	// list of card suits
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	// list of card values
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// for loops iterate through each one
	// for every combination of suit and value to add a new card to our cards
	// _ to ignore unused variable (index)
	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value+" of "+suit)
		}
	}

	// return slice (deck of cards)
	return cards
}

// loop through deck of cards
// prints out the value
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// get two decks (hand of cards, remaining cards)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// turn deck into a string
func (d deck) toString() string {
	// strings join takes our slice of string
	// join it all into one single string separated by commas
	return strings.Join([]string(d), ",")
}