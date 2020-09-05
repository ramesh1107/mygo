package main

import (
	"fmt"

)
func main (){

	//print()
	numbr := [] int{0,1,2,3,4,5,6,7,8,9,10}
	//var k int =0
	//fmt.Println(numbr)

	for i, num:= range numbr  {
		num = numbr[i]

	if num % 2 == 0{

		fmt.Println(num, " is even", )
	} else {
			fmt.Println(num, " is odd",)
		}

	}
}

