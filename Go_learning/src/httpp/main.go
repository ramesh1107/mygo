package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		log.Println("Hey ramesh")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "helo %s", d)

	})
	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("Hey SRINIDHI")
	})
	http.ListenAndServe(":9090", nil)
}
