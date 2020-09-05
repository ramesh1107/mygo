package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main(){

	file, err := os.Open(file.txt) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	bs := make([]byte, 100)
	count, err := file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, bs[:count])


}



