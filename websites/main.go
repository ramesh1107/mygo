package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{

		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checksite(link, c)

	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checksite(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {

		fmt.Println(link, "is down ")
		c <- "link is down"

		return
	}

	fmt.Println(link, "is up and running")

	c <- "link is up"

}
