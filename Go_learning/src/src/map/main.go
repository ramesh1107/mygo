package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#008000",
		"white": "#ff80f0",
	}
	printMap(colors)
}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)

	}
}
