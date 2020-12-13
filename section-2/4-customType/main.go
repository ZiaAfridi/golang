package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "New Value")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
