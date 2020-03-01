package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/andrewmathies/msl/data"
	"github.com/gorilla/mux"
)

type ERDs struct {
	logger *log.Logger
}

type KeyERD struct{}

func NewERDs(l *log.Logger) *ERDs {
	return &ERDs{l}
}

func (e *ERDs) GetERDs(rw http.ResponseWriter, r *http.Request) {
	e.logger.Println("Handle GET ERDs")

	erds := data.GetERDs()
	err := erds.ToJSON(rw)

	if err != nil {
		e.logger.Println(err)
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (e *ERDs) AddERD(rw http.ResponseWriter, r *http.Request) {
	e.logger.Println("Handle POST ERD")

	erd := r.Context().Value(KeyERD{}).(data.ERD)

	e.logger.Printf("ERD: %#v", erd)
	data.AddERD(&erd)
}

func (e *ERDs) UpdateERD(rw http.ResponseWriter, r *http.Request) {
	e.logger.Println("Handle PUT ERD")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		e.logger.Println(err)
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
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

func (e ERDs) MiddlewareValidateERD(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		erd := data.ERD{}
		err := erd.FromJSON(r.Body)

		if err != nil {
			e.logger.Println(err)
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// validate the erd
		err = erd.Validate()

		if err != nil {
			e.logger.Println("ERROR validating product", err)
			http.Error(rw, fmt.Sprintf("Invalid request data: %s", err), http.StatusBadRequest)
			return
		}

		// add the erd to the context
		ctx := context.WithValue(r.Context(), KeyERD{}, erd)
		r = r.WithContext(ctx)

		// call the next handler. could be another middleware or the final handler
		next.ServeHTTP(rw, r)
	})
}
