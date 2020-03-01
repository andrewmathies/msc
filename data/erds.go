package data

import (
	"encoding/json"
	"fmt"
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

// dummy data until we get to db stuff
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

func getNextID() int {
	lastERD := erds[len(erds)-1]
	return lastERD.ID + 1
}

func AddERD(erd *ERD) {
	erd.ID = getNextID()
	erds = append(erds, erd)
}

func UpdateERD(id int, erd *ERD) error {
	_, i, err := findERD(id)

	if err != nil {
		return err
	}

	erd.ID = id
	erds[i] = erd

	return nil
}

var ErrERDNotFound = fmt.Errorf("ERD not found")

func findERD(id int) (*ERD, int, error) {
	for i, erd := range erds {
		if erd.ID == id {
			return erd, i, nil
		}
	}

	return nil, -1, ErrERDNotFound
}

func GetERDs() ERDs {
	return erds
}

func (e *ERDs) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(e)
}

func (e *ERD) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(e)
}
