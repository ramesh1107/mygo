package main

import "fmt"
type contactInfo struct{

	email string
	zipcode int
}
type person struct{

	firstName string
	lastName string
	contact contactInfo
}
func main (){

	jim:= person{firstName:"jim",
		lastName:"anderson",
		contact : contactInfo{
			email:"a@b.com",
			zipcode:12345,
		},

	}
	jim.print()


}
  func (p person) print(){
	  fmt.Printf("%+v", p)
  }