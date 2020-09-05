package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID        int
	Name      string
	Desc      string
	Price     float32
	SKU       string
	Createdon string
	UpdatedOn string
	DeletedOn string
}

func (p *Product) FromJSON(r io.Reader) error {

	e := json.NewDecoder(r)
	return e.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {

	e := json.NewEncoder(w)
	return e.Encode(p)
}
func GetProducts() Products {
	return productList
}
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdatedProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil

}

// error def

var ErrPNotfound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {

			return p, i, nil
		}
	}
	return nil, -1, ErrPNotfound
}
func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
	&Product{
		ID:        1,
		Name:      "Latte",
		Desc:      "string",
		Price:     2.50,
		SKU:       "abc234",
		Createdon: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID:        2,
		Name:      "Capochino",
		Desc:      "string",
		Price:     3.50,
		SKU:       "abcd1234",
		Createdon: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID:        3,
		Name:      "Expresso",
		Desc:      "string",
		Price:     1.50,
		SKU:       "ab123456",
		Createdon: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
