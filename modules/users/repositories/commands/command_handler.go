package commands

import "net/http"

// UsersCommandHandler Interface ..
type UsersCommandHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request) error
	Login(w http.ResponseWriter, r *http.Request) error
}
