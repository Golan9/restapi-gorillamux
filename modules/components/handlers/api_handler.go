package handlers

import (
	"log"
	"net/http"
	cfg "restapi-gorillamux/config"
	com "restapi-gorillamux/modules/components/repositories/commands"
	q "restapi-gorillamux/modules/components/repositories/queries"
)

var config = cfg.Config{}

// APIHandler ..
type APIHandler struct {
	ch com.ComponentsCommandHandler
	qh q.ComponentsQueryHandler
}

// NewCommands .. // DITAMBAHIN RAHMAT
func NewCommands() com.ComponentsCommandHandler {
	return &com.ComponentsCommandHandlers{}
}

// NewUsers ..
func NewUsers() q.ComponentsQueryHandler {
	return &q.ComponentsQueryHandlers{}
}

// APICommandHandler ..
type APICommandHandler interface {
	InsertCompHandler(w http.ResponseWriter, r *http.Request)
	UpdateCompHandler(w http.ResponseWriter, r *http.Request)
	DeleteCompHandler(w http.ResponseWriter, r *http.Request)
	GetAllCompHandler(w http.ResponseWriter, r *http.Request)
	GetOneCompHandler(w http.ResponseWriter, r *http.Request)
}

// NewAPIHandler .. // RAHMAT JUGA
func NewAPIHandler(ch com.ComponentsCommandHandler, qh q.ComponentsQueryHandler) APICommandHandler {
	return &APIHandler{ch: ch, qh: qh}
}

// InsertCompHandler ..
func (api *APIHandler) InsertCompHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("cek error1")

	returnVals := api.ch.InsertComponent(w, r)

	//fmt.Println("cek error", returnVals)

	if returnVals != nil {
		//fmt.Println("error disini cuy")
		log.Println("Something went wrong")
		return
	}
	//fmt.Println("masuk disini cuy")

	return
}

// UpdateCompHandler ..
func (api *APIHandler) UpdateCompHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("cek error1")

	returnVals := api.ch.UpdateComponent(w, r)

	//fmt.Println("cek error", returnVals)

	if returnVals != nil {
		//fmt.Println("error disini cuy")
		log.Println("Something went wrong")
		return
	}
	//fmt.Println("masuk disini cuy")

	return
}

// DeleteCompHandler ..
func (api *APIHandler) DeleteCompHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("cek error1")

	returnVals := api.ch.DeleteComponent(w, r)

	//fmt.Println("cek error", returnVals)

	if returnVals != nil {
		//fmt.Println("error disini cuy")
		log.Println("Something went wrong")
		return
	}
	//fmt.Println("masuk disini cuy")

	return
}

// GetAllCompHandler ..
func (api *APIHandler) GetAllCompHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("cek error1")

	returnVals := api.qh.GetAllComponent(w, r)

	//fmt.Println("cek error", returnVals)

	if returnVals != nil {
		//fmt.Println("error disini cuy")
		log.Println("Something went wrong")
		return
	}
	//fmt.Println("masuk disini cuy")

	return
}

// GetOneCompHandler ..
func (api *APIHandler) GetOneCompHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("cek error1")

	returnVals := api.qh.GetOneComponent(w, r)

	//fmt.Println("cek error", returnVals)

	if returnVals != nil {
		//fmt.Println("error disini cuy")
		log.Println("Something went wrong")
		return
	}
	//fmt.Println("masuk disini cuy")

	return
}
