package handlers

import (
	"log"
	"net/http"

	"github.com/andrewmathies/msl/data"
)

type ERDs struct {
	logger *log.Logger
}

func NewERDs(l *log.Logger) *ERDs {
	return &ERDs{l}
}

func (e *ERDs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		e.getERDs(rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (e *ERDs) getERDs(rw http.ResponseWriter, r *http.Request) {
	erds := data.GetERDs()
	err := erds.ToJSON(rw)

	if err != nil {
		e.logger.Println(err)
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
