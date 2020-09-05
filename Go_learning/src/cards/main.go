package main

func main() {

	//cards := newDeckFromFile("my_cards")
	//hand, reaminingCards := deal(cards, 5)
	cards := newDeck()
	//cards.print()
	//hand.print()
	//reaminingCards.print()

	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	cards.shuffle()
	cards.print()
}
