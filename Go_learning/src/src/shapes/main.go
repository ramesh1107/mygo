package main

import "fmt"

type shape interface {
	findArea() float64
}
type square struct {
	sideLength float64
}
type traingle struct {
	base   float64
	height float64
}

func main() {

	fmt.Println("Boiler plate")

	s := square{sideLength: 10}
	t := traingle{base: 10, height: 10}
	
	printArea(s)
	printArea(t)

}

func (s square) findArea() float64 {
	
	return s.sideLength * s.sideLength

}
func (t traingle) findArea() float64 {
	
	return 0.5 * t.base * t.height

}

func printArea(s shape) {

	fmt.Println("Area of the shape is", s.findArea())

}
