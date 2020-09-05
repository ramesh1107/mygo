package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type lW struct{}

func main() {

	resp, err := http.Get("http://news.google.com")

	if err != nil {

		fmt.Println("Error:", err)

		os.Exit(1)
	}

	//bs := make([]byte, 999999)
	//resp.Body.Read(bs)

	//fmt.Println("withouth sting", bs)
	l := lW{}
	io.Copy(l, resp.Body)

}

func (lW) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("Just wrorte bytes", len(bs))
	return len(bs), nil

}
