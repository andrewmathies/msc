package handlers

import (
	"log"
	"net/http"

	"github.com/andrewmathies/msl/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)

	if err != nil {
		p.logger.Panic(err)
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
