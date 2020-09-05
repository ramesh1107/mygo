package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{

		"http://google.com",
		"http://facebook.com",
		"http://news.google.com",
		"http://razorpay.com",
		"http://flipkart.in",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}

	for i := 0; i < len(links); i++ {

		go checkLink(<-c, c)
	}
}
func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return

	}

	fmt.Println(link, "is working")
	c <- link
}
