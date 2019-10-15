package queries

import "net/http"

// UsersQueryHandler Interface ..
type UsersQueryHandler interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request) error
}
