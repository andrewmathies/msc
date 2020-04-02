package handlers

import (
	"net/http"

	"github.com/andrewmathies/msl/data"
)

func (e *ERDs) AddERD(rw http.ResponseWriter, r *http.Request) {
	e.logger.Println("Handle POST ERD")

	erd := r.Context().Value(KeyERD{}).(data.ERD)

	e.logger.Printf("ERD: %#v", erd)
	data.AddERD(&erd)
}
