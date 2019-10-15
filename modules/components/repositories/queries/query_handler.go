package queries

import "net/http"

// ComponentsQueryHandler Interface ..
type ComponentsQueryHandler interface {
	GetAllComponent(w http.ResponseWriter, r *http.Request) error
	GetOneComponent(w http.ResponseWriter, r *http.Request) error
}
