package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	if r.Method == http.MethodPut {
		regex := regexp.MustCompile("/([0-9]+)")
		matches := regex.FindAllStringSubmatch(r.URL.Path, -1)

		if len(matches) != 1 {
			p.l.Println("Invalid URI: more than 1 ids")
			http.Error(w, "More than 1 ids provided", http.StatusBadRequest)
			return
		}
		if len(matches[0]) != 2 {
			p.l.Println("Invalid URI: More than 2 groups")
			http.Error(w, "More than 2 capture groups", http.StatusBadRequest)
			return
		}
		idStr := matches[0][1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			p.l.Println("Unable to convert id string to id")
			http.Error(w, "Invalid id provided", http.StatusBadRequest)
		}

		p.updateProduct(id, w, r)
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
	w.Header().Add("Content-Type", "application/json")
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

func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT request")

	//convert method body to Product type
	newProduct := &data.Product{}
	err := newProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, newProduct)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusBadRequest)
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
	}
}
