// Package classification ERD API
//
// Documentation for ERD API
//
// 	Schemes: http
// 	BasePath: /
// 	Version: 0.0.1
//
// 	Consumes:
// 	- application/json
//
// 	Produces:
// 	- application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/andrewmathies/msl/data"
)

// A list of ERDs in the response
// swagger:response erdsResponse
type erdsResponse struct {
	// All erds in the system
	// in:body
	Body []data.ERD
}

// swagger:response noContent
type erdsNoContent struct {
}

// swagger:parameters deleteERD
type erdIDParameterWrapper struct {
	// The id of the ERD to remove from the data store
	// in: path
	// required: true
	ID int `json:"id"`
}

type ERDs struct {
	logger *log.Logger
}

type KeyERD struct{}

func NewERDs(l *log.Logger) *ERDs {
	return &ERDs{l}
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
