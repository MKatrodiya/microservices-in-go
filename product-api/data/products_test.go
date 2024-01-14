package data

import "testing"

func TestProductValidations(t *testing.T) {
	p := &Product{
		Name:  "Water",
		Price: 1.00,
		SKU:   "abc-def-ghijkl",
	}

	err := p.ValidateProduct()

	if err != nil {
		t.Fatal(err)
	}
}
