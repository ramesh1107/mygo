package main

import "fmt"

type englishbot struct{}
type spanishbot struct{}
type bot interface {
	getGreeting() string
}

func main (){

	eb:= englishbot{}
	sb:= spanishbot{}
	printgreeting(eb)
	printgreeting(sb)


}

func (englishbot) getGreeting() string {

	return "Hello Kris"


}

func (spanishbot) getGreeting() string {

	return "f Kris"

}

func printgreeting (b bot){

	fmt.Println(b.getGreeting())
}

