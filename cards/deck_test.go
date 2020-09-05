package main

import "testing"

func TestnewDeck(t * testing.T) {

	d := newDeck()

	if len(d)!= 2 {
		t.Errorf ("expected deck of lenght 52 but got %v", len(d) )
	}

	if d[0] != "Ace of spades"{

		t.Errorf ("expected first card to be Ace of spades but got %v", d[0] )
}

	if d[len(d)-1 ] != "Ace of clubs" {

		t.Errorf ("expected first card to be four of clubs but got %v", d[len(d)-1] )
	}
}