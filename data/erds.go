package data

import "net/http"

type ERD struct {
	name    string
	address uint16
	fields  []string
	actions []string
}

type ERDs *[]ERD

var erds *[]ERD

func GetERDs() ERDs {
	return erds
}

/*
func LoadERDs(path string) {
	return
}
*/
func ToJSON(rw http.ResponseWriter) error {

}
