package main
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := [] string {"spades", "daimonds","hearts","clubs"}
	cardValues := [] string {"Ace", "two","three","four"}
							//"five", "six","seven","eight",
							//"nine", "jack","queen","king"}

	for _, suit:= range cardSuits  {

		for _ ,value := range cardValues {

			cards = append (cards, value + " of " +suit)
		}
	}
return cards
}

func (d deck) print() {
	for i, card := range d {

		fmt.Println(i,card)
	}
}

func deal (d deck, handsize int) (deck, deck){
	return  d[:handsize], d[handsize:]

}

func (d deck) toString() string {
	return strings.Join([]string (d), ",")
}

func (d deck) sFT (fname string) error {
return	ioutil.WriteFile(fname, []byte(d.toString()),777)


	}

func newDeckFromFile(filename string) deck{
	bs, err:= ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	s:= strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle(){

	for i := range d{

		nP := rand.Intn(len(d) -1)

		d[i], d[nP]= d[nP], d[i]
	}

}