package handlers

import (
	"net/http"

	"github.com/andrewmathies/msl/data"
)

// swagger:route GET /erds erds listERDs
// Returns a list of erds
// responses:
// 	200: erdsResponse

// GetERDs returns the erds from the data store
func (e *ERDs) GetERDs(rw http.ResponseWriter, r *http.Request) {
	e.logger.Println("Handle GET ERDs")

	// fetch erds from data store
	erds := data.GetERDs()

	// serialize erds to JSON
	err := erds.ToJSON(rw)

	if err != nil {
		e.logger.Println(err)
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
