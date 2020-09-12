// First go handler for Crud operations
//
// Documentation for this API
//
// schemes:
//  - https
// basePath: /
//
//
//consumes:
//  - application/json
// - application/xml
//produces:
//  - application/json
//  - application/xml
// swagger:meta

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

// A list of products returns in the response
//swagger:response productsResponse

type productsResponseWrapper struct {
	//All the products in the system
	// in: body
	Body []data.Product
}
type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {

	return &Products{l}
}

// swagger:route GET /products products listProducts
// Returns a list of Products
// responses:
// 200:productsResposne

// GetPorducts returns products from data store

func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle Get Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {

		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route POST /products products addProducts
// Adds a new Product to the product lists

// AddProduct adds a new products to the data store
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle addition of Products")

	//prod := &data.Product{}
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
	//err := prod.FromJSON(r.Body)
	/*if err != nil {
		http.Error(rw, "unable to unmarshal json- add", http.StatusBadRequest)
	}*/
}

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
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

// MiddlewareValidateProduct validates the product in the request and calls next if ok
func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

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
				fmt.Sprintf("Error validating Product: %s", err),
				http.StatusBadRequest)
			return

		}

		//ctx :=r.Context().WithValue(KeyProduct{},prod)
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
