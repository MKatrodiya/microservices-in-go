package data

import (
	"testing"
)

func TestProductValidations(t *testing.T) {
	p := &Product{
		Name:  "Water",
		Price: 0.00,
		SKU:   "abc-def-ghijkl",
	}

	validator := NewValidation()
	err := validator.Validate(p)

	if err != nil {
		t.Fatal(err)
	}
}
