package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MKatrodiya/ProductMicroservices/data"
)

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
		errs := p.v.Validate(prod)
		if errs != nil {
			p.l.Println("[ERROR]: Validating product data")
			http.Error(
				w,
				fmt.Sprintf("Invalid product data: %s", errs),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
