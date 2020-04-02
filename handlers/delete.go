package handlers

import (
	"net/http"
	"strconv"

	"github.com/andrewmathies/msl/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /erds/{id} erds deleteERD
// Delete an ERD
// responses:
// 	201: noContent

// DeleteERD deletes a product from the data store
func (e *ERDs) DeleteERD(rw http.ResponseWriter, r *http.Request) {
	e.logger.Println("Handle DELETE ERD")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		e.logger.Println(err)
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	err = data.DeleteERD(id)

	if err == data.ErrERDNotFound {
		http.Error(rw, "ERD not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "ERD not found", http.StatusInternalServerError)
		return
	}
}
