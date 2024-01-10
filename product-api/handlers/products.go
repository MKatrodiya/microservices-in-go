package handlers

import (
	"log"
	"net/http"

	"github.com/MKatrodiya/ProductMicroservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	//all other methods
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	productsList := data.GetProducts()

	err := productsList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal data to json", http.StatusInternalServerError)
	}
}
