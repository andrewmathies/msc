package data

import (
	"encoding/json"
	"io"
	"time"
)

type ERD struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Address   string   `json:"address"`
	Fields    []string `json:"fields"`
	Actions   []string `json:"actions"`
	CreatedOn string   `json:"-"`
	UpdatedOn string   `json:"-"`
	DeletedOn string   `json:"-"`
}

type ERDs []*ERD

var erds = []*ERD{
	&ERD{
		ID:        0,
		Name:      "ERD_ApplianceType",
		Address:   "0x0090",
		Fields:    []string{"ApplianceType"},
		Actions:   []string{"Read", "Subscribe"},
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&ERD{
		ID:        1,
		Name:      "ERD_OvenInfo",
		Address:   "0x5007",
		Fields:    []string{"OvenInfo"},
		Actions:   []string{"Read", "Write", "Subscribe", "Publish"},
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}

func GetERDs() ERDs {
	return erds
}

func (e *ERDs) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(e)
}
