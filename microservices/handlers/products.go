package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mygo/microservices/data"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {

	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle Get Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {

		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle addition of Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {

		http.Error(rw, "unable to unmarshal json- add", http.StatusBadRequest)

	}
	data.AddProduct(prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	// Adding Gorilla package code
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Error in id string conversion", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle Update Product", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdatedProduct(id, &prod)
	if err == data.ErrPNotfound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return

	}

	if err != nil {
		http.Error(rw, "product not found", http.StatusInternalServerError)
		return

	}

}

type KeyProduct struct{}

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserilizing product")
			http.Error(rw, "Error Reading Product", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()

		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(rw,
				fmt.Sprintln("Error validating Product: %s", err),
				http.StatusBadRequest)
			return

		}

		//ctx :=r.Context().WithValue(KeyProduct{},prod)
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
