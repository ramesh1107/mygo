package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Please give one argument.")
		return
	}

	filename := os.Args[1]
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error in opening ", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	io.Copy(os.Stdout, f)

}
