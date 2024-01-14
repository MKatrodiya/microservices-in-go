package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Product) ValidateProduct() error {
	validator := validator.New()
	validator.RegisterValidation("sku", validateSKU)
	return validator.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	//format of SKU: abc-def-vjhvgj
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) == 1 {
		return true
	}

	return false
}

func GetProducts() Products {
	return productsList
}

func AddProduct(p *Product) {
	p.ID = getNextID()

	productsList = append(productsList, p)
}

func getNextID() int {
	lp := productsList[len(productsList)-1]
	return lp.ID + 1
}

func UpdateProduct(id int, updatedProduct *Product) error {
	_, pos, err := getProductByID(id)
	if err != nil {
		return err
	}
	updatedProduct.ID = id
	productsList[pos] = updatedProduct
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

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
