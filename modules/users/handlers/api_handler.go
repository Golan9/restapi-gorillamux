package handlers

import (
	"log"
	"net/http"
	cfg "restapi-gorillamux/config"
	com "restapi-gorillamux/modules/users/repositories/commands"
	q "restapi-gorillamux/modules/users/repositories/queries"
)

var config = cfg.Config{}

// APIHandler ..
type APIHandler struct {
	ch com.UsersCommandHandler
	qh q.UsersQueryHandler
}

// NewCommands .. // DITAMBAHIN RAHMAT
func NewCommands() com.UsersCommandHandler {
	return &com.UsersCommandHandlers{}
}

// NewUsers ..
func NewUsers() q.UsersQueryHandler {
	return &q.UsersQueryHandlers{}
}

// APICommandHandler ..
type APICommandHandler interface {
	RegisterUserHandler(w http.ResponseWriter, r *http.Request)
	PostDataLogin(w http.ResponseWriter, r *http.Request)
	GetUsersHandler(w http.ResponseWriter, r *http.Request)
}

// NewAPIHandler .. // RAHMAT JUGA
func NewAPIHandler(ch com.UsersCommandHandler, qh q.UsersQueryHandler) APICommandHandler {
	return &APIHandler{ch: ch, qh: qh}
}

// RegisterUserHandler ..
func (api *APIHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("cek error1")

	returnVals := api.ch.RegisterUser(w, r)

	//fmt.Println("cek error", returnVals)

	if returnVals != nil {
		//fmt.Println("error disini cuy")
		log.Println("Something went wrong")
		return
	}
	//fmt.Println("masuk disini cuy")

	return
}

// PostDataLogin ..
func (api *APIHandler) PostDataLogin(w http.ResponseWriter, r *http.Request) {

	returnVals := api.ch.Login(w, r)

	if returnVals != nil {
		log.Println("Something went wrong")
		return
	}
	return
}

// GetUsersHandler ..
func (api *APIHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	returnVals := api.qh.GetAllUsers(w, r)

	if returnVals != nil {
		log.Println("Something went wrong")
		return
	}
	return
}
