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
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	//all other methods
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request")
	//fetch products from the datastore
	productsList := data.GetProducts()

	//serialize the list to JSON
	err := productsList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal data to json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST request")

	//convert method body to Product type
	newProduct := &data.Product{}
	err := newProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
	}
	data.AddProduct(newProduct)
}
