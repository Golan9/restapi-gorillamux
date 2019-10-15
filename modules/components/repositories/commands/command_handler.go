package commands

import "net/http"

// ComponentsCommandHandler Interface ..
type ComponentsCommandHandler interface {
	InsertComponent(w http.ResponseWriter, r *http.Request) error
	UpdateComponent(w http.ResponseWriter, r *http.Request) error
	DeleteComponent(w http.ResponseWriter, r *http.Request) error
}
