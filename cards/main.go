package main



func main (){
	cards  :=newDeckFromFile("kris")
	//hand, reamainingCards:= deal(cards,5)
	cards.shuffle()
	cards.print()
	//reamainingCards.print()
	//fmt.Print(cards.toString())


}

