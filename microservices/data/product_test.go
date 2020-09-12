package data

import "testing"

func TestValidations(t *testing.T) {

	p := &Product{

		Name:  "ramesh",
		Price: 1.00,
		SKU:   "abc-def-ghk",
	}

	err := p.Validate()

	if err != nil {

		t.Fatal(err)
	}
}
