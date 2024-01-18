package handlers

import (
	"net/http"

	"github.com/MKatrodiya/ProductMicroservices/data"
)

// swagger:route PUT /products products updateProduct
// Update a products
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) Update(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT request")
	id := getProductID(r)

	//fetxh product from the context
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err := data.UpdateProduct(id, *prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in the database"}, w)
		return
	}

	// write the no content success header
	w.WriteHeader(http.StatusNoContent)
}
