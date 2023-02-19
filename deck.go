package main

import "fmt"

// create a new type of "deck"
// which is a slice of strings

type deck []string

func (deneme deck) print() {
	for i, card := range deneme {
		fmt.Println(i, card)
	}
}
