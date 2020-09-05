package main

import "fmt"

type bot interface {
	getGreeting() string
}
type engBot struct {
}
type spnBot struct {
}

func main() {

	eb := engBot{}
	sb := spnBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (engBot) getGreeting() string {

	return "hello!!!"

}
func (spnBot) getGreeting() string {

	return "hola!!!"

}

func printGreeting(b bot) {

	fmt.Println(b.getGreeting())

}
