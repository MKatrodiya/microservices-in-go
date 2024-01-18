package handlers

import (
	"net/http"

	"github.com/MKatrodiya/ProductMicroservices/data"
)

// swagger:route GET /products products getAll
// Returns the list of products
// responses:
// 		200:productsResponse

// GetAll returns the list of products from the data store
func (p *Products) GetAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request")
	//fetch products from the datastore
	productsList := data.GetProducts()

	//serialize the list to JSON
	w.Header().Add("Content-Type", "application/json")
	err := data.ToJSON(productsList, w)
	if err != nil {
		p.l.Println("[ERROR] serializing products", err)
	}
}

// swagger:route GET /products/{id} products getSingle
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// GetSingle returns the product with id given in request path
func (p *Products) GetSingle(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return

	default:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = data.ToJSON(prod, w)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}
