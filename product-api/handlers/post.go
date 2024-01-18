package handlers

import (
	"net/http"

	"github.com/MKatrodiya/ProductMicroservices/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//	422: errorValidation
//	501: errorResponse

// Add handles POST requests and asdds a new product to the database
func (p *Products) Add(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST request")

	//fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(*prod)
}
