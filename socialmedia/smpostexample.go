package main

import (
	"fmt"
	"socialmedia"
)

func main(){

mypost := socialmedia.NewPost("Ramesh",socialmedia.Moods["thrilled"], "Go is Awsome!!", "Check out the website", "https://Golang.org",
	"","",[]string{"go","golang","pgm lang"})

fmt.Printf("mypost :%+v\n",mypost)

}
