package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/mygo/microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {

	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("Putted", r.URL.Path)
		reg := regexp.MustCompile("/([0-9]+)")
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("Invalid more than one id ")
			http.Error(rw, "Invalid Uri", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid more than one capture group ")
			http.Error(rw, "Invalid Uri", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		p.l.Println("Invalid URL unable to convert to number", idString)
		if err != nil {
			http.Error(rw, "Invalid Uri", http.StatusBadRequest)
			return
		}
		p.updateProducts(id, rw, r)

		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)

}
func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {

		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Product")
	//lp := data.GetProducts()
	//err := lp.ToJSON(rw)

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {

		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
		data.AddProduct(prod)
	}
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Post Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {

		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)

	}
	err = data.UpdatedProduct(id, prod)
	if err == data.ErrPNotfound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return

	}

	if err != nil {
		http.Error(rw, "product not found", http.StatusInternalServerError)
		return

	}

}
