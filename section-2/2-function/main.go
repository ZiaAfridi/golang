package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

// golang function and return type of function.
func newCard() string {
	return "Ace of Diamond"
}
