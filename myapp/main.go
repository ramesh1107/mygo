package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

const usname string = `^@?(\w){1.15}$`
func main(){

	var uname  string

	flag.StringVar(&uname, "Username","great","Function to check name")
	flag.Parse()


	fmt.Println("Fucntion to check user name")
	fmt.Println("checking syntax for username, \"", uname, "\", resulted in:  ", Checkuname(uname), "\n")

}
func Checkuname(username string)  bool{

	validationResult :=false
	r, err:= regexp.Compile(usname)
	if err !=nil	{
		log.Fatal(err)
	}
	validationResult= r.MatchString(usname)
	return validationResult
}