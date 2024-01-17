package data

import (
	"fmt"
	"time"
)

// Product defines a structure for an API product
// swagger: model
type Product struct {
	// the id of the product
	//
	// required: false
	// min: 1
	ID int `json:"id"`
	// the name of the product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`
	// the description of the product
	//
	// required; false
	// max length: 10000
	Description string `json:"description"`
	// the price of the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`
	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU       string `json:"sku" validate:"required,sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// Products defines a slice of the products
type Products []*Product

// GetProducts returns all products from the database
func GetProducts() Products {
	return productsList
}

// AddProduct adds a new product to the database
func AddProduct(p *Product) {
	p.ID = getNextID()

	productsList = append(productsList, p)
}

// getNextID returns the next available id for the product
func getNextID() int {
	lp := productsList[len(productsList)-1]
	return lp.ID + 1
}

// UpdateProduct modifies the existing product with the given id
// with the given updated product
// If a product is not found, this returns a ProductNotFound error
func UpdateProduct(id int, updatedProduct *Product) error {
	_, pos, err := getProductByID(id)
	if err != nil {
		return err
	}
	updatedProduct.ID = id
	productsList[pos] = updatedProduct
	return nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productsList = append(productsList[:i], productsList[i+1])

	return nil
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productsList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

// ErrProductNotFound error is raised when a product is not found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

// getProductByID returns a product from the database whose id matches with
// the given id
// If a product is not found, this returns a ProductNotFound error
func getProductByID(id int) (*Product, int, error) {
	for i, p := range productsList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var productsList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
