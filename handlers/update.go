package handlers

import (
	"net/http"
	"strconv"

	"github.com/andrewmathies/msl/data"
	"github.com/gorilla/mux"
)

func (e *ERDs) UpdateERD(rw http.ResponseWriter, r *http.Request) {
	e.logger.Println("Handle PUT ERD")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		e.logger.Println(err)
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	erd := r.Context().Value(KeyERD{}).(data.ERD)

	e.logger.Printf("ERD: %#v", erd)
	err = data.UpdateERD(id, &erd)

	if err == data.ErrERDNotFound {
		http.Error(rw, "ERD not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Unable to update ERD", http.StatusInternalServerError)
		return
	}
}
