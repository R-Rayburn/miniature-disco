package main

//func main() {
//	// var card string = "Ace of Spades"
//	// var card = "Ace of Spades"
//	// var card string
//	// card = "Ace of Spades"
//	card := newCard() // "Ace of Spades" // Only for new variables (initialization), not re-assignment
//	// card = "Five of Diamonds"
//
//	fmt.Println(card)
//
//	cards := []string{
//		newCard(),
//		"Ace of Diamonds",
//		newCard(),
//	}
//	cards = append(cards, "Six of Spades")
//
//	for i, card := range cards {
//		fmt.Println(i, card)
//	}
//}

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards = newDeck(false)
	//cards.print()

	hand, cards := deal(cards, 5)

	hand.print()
	cards.print()
	cards.saveToFile("filename.txt")

	deckFromFile := newDeckFromFile("filename.txt")
	deckFromFile.shuffle()
	deckFromFile.print()
}
