package main

import "fmt"

type person struct {
	fname string
	lname string
	contactInfo
}

type contactInfo struct {
	email string
	pincd int
}

func main() {

	alex := person{
		fname: "ramesh",
		lname: "a",
		contactInfo: contactInfo{
			email: "ramesh@gmail.com",
			pincd: 560035,
		},
	}
	//alexPointer := &alex
	alex.updateName("akkanapragada")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newlname string) {
	(*pointerToPerson).lname = newlname
}
