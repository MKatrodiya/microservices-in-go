// Package classification of Product API
//
// # Documentation for Product API
//
// Schemes: HTTP
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MKatrodiya/ProductMicroservices/data"
	"github.com/gorilla/mux"
)

// A list of products returned in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All the products in system
	// in: body
	Body []*data.Product
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /products products listProducts
// Returns the list of products
// responses:
// 		200:productsResponse

// GetProducts returns the list of products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request")
	//fetch products from the datastore
	productsList := data.GetProducts()

	//serialize the list to JSON
	w.Header().Add("Content-Type", "application/json")
	err := data.ToJSON(productsList, w)
	if err != nil {
		http.Error(w, "Unable to marshal data to json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST request")

	//convert method body to Product type
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT request")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
	}

	//convert method body to Product type
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Serialize product from request body
		prod := data.Product{}
		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR]: Deserializing product data")
			http.Error(w, "Invalid product data", http.StatusBadRequest)
			return
		}

		//Validate product
		err = prod.ValidateProduct()
		if err != nil {
			p.l.Println("[ERROR]: Validating product data")
			http.Error(
				w,
				fmt.Sprintf("Invalid product data: %s", err),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
