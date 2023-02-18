package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades", "ibram")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)

}

func newCard() string {
	return "Five of Diamonds"
}
