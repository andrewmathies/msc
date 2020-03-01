package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/andrewmathies/msl/data"
)

type ERDs struct {
	logger *log.Logger
}

func NewERDs(l *log.Logger) *ERDs {
	return &ERDs{l}
}

// ServeHTTP implements the go http.Handler interface
func (e *ERDs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		e.getERDs(rw, r)
	case http.MethodPost:
		e.addERD(rw, r)
	case http.MethodPut:
		// expect the id in the URI
		re := regexp.MustCompile(`([0-9]+)`)
		group := re.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 || len(group[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			e.logger.Println(err)
			http.Error(rw, "couldn't parse id", http.StatusInternalServerError)
			return
		}

		e.logger.Println("got id", id)

		e.updateERD(id, rw, r)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (e *ERDs) getERDs(rw http.ResponseWriter, r *http.Request) {
	e.logger.Println("Handle GET ERDs")

	erds := data.GetERDs()
	err := erds.ToJSON(rw)

	if err != nil {
		e.logger.Println(err)
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (e *ERDs) addERD(rw http.ResponseWriter, r *http.Request) {
	e.logger.Println("Handle POST ERD")

	erd := &data.ERD{}
	err := erd.FromJSON(r.Body)

	if err != nil {
		e.logger.Println(err)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	e.logger.Printf("ERD: %#v", erd)
	data.AddERD(erd)
}

func (e *ERDs) updateERD(id int, rw http.ResponseWriter, r *http.Request) {
	e.logger.Println("Handle PUT ERD")

	erd := &data.ERD{}
	err := erd.FromJSON(r.Body)

	if err != nil {
		e.logger.Println(err)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	e.logger.Printf("ERD: %#v", erd)
	err = data.UpdateERD(id, erd)

	if err == data.ErrERDNotFound {
		http.Error(rw, "ERD not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(rw, "Unable to update ERD", http.StatusInternalServerError)
		return
	}
}
