package main

import (
	"testing"
)

func TestNeWDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 100 {

		t.Errorf("Expected deck lenght of 16 but got %v ", len(d))
	}

}
