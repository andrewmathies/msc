package data

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type ERD struct {
	ID        int      `json:"id"`
	Name      string   `json:"name" validate:"required"`
	Address   string   `json:"address" validate:"required"`
	Fields    []string `json:"fields" validate:"gt=0,dive,required,validerdfield=Name"`
	Actions   []string `json:"actions"`
	CreatedOn string   `json:"-"`
	UpdatedOn string   `json:"-"`
	DeletedOn string   `json:"-"`
}

type ERDs []*ERD

// Custom errors
var ErrERDNotFound = fmt.Errorf("ERD not found")

// dummy data until we get to db stuff
var erds = []*ERD{
	&ERD{
		ID:        0,
		Name:      "ERD_ApplianceType",
		Address:   "0x0090",
		Fields:    []string{"ERD_ApplianceType"},
		Actions:   []string{"Read", "Subscribe"},
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&ERD{
		ID:        1,
		Name:      "ERD_OvenInfo",
		Address:   "0x5007",
		Fields:    []string{"ERD_OvenInfo"},
		Actions:   []string{"Read", "Write", "Subscribe", "Publish"},
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}

func getNextID() int {
	lastERD := erds[len(erds)-1]
	return lastERD.ID + 1
}

// CRUD functions for in-memory list acting as db
func GetERDs() ERDs {
	return erds
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

func findERD(id int) (*ERD, int, error) {
	for i, erd := range erds {
		if erd.ID == id {
			return erd, i, nil
		}
	}

	return nil, -1, ErrERDNotFound
}

// JSON serialization/deserialization
func (e *ERDs) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(e)
}

func (e *ERD) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(e)
}

// Validation
func (e *ERD) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("validerdfield", validateFields)
	return validate.Struct(e)
}

// this is checking that all the names of all ERD fields contain the ERD name itself
func validateFields(fl validator.FieldLevel) bool {
	name, _, ok := fl.GetStructFieldOK()

	if !ok {
		return false
	}

	erdField := fl.Field()

	return strings.Contains(erdField.String(), name.String())
}
