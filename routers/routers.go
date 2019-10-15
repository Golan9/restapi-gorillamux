package routers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	cfg "restapi-gorillamux/config"
	database "restapi-gorillamux/database"
	auth "restapi-gorillamux/jwt"
	comps "restapi-gorillamux/modules/components/handlers"
	users "restapi-gorillamux/modules/users/handlers"
)

var dao = database.DAO{}
var config = cfg.Config{}

// Init ..
func Init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Routers ..
func Routers() {

	// init Routers
	Router := mux.NewRouter()

	// init handlers

	// users ..
	userCom := users.NewCommands()
	userQ := users.NewUsers()
	usersHandler := users.NewAPIHandler(userCom, userQ)

	// components ..
	compCom := comps.NewCommands()
	compQ := comps.NewUsers()
	compsHandler := comps.NewAPIHandler(compCom, compQ)

	// Api Handler / Endpoints

	// users ..
	var registerUserHandler = http.HandlerFunc(usersHandler.RegisterUserHandler)
	var postDataLogin = http.HandlerFunc(usersHandler.PostDataLogin)
	var getUsersHandler = http.HandlerFunc(usersHandler.GetUsersHandler)

	// components ..
	var insertCompHandler = http.HandlerFunc(compsHandler.InsertCompHandler)
	var getAllCompHandler = http.HandlerFunc(compsHandler.GetAllCompHandler)
	var getOneCompHandler = http.HandlerFunc(compsHandler.GetOneCompHandler)
	var updateCompHandler = http.HandlerFunc(compsHandler.UpdateCompHandler)
	var deleteCompHandler = http.HandlerFunc(compsHandler.DeleteCompHandler)

	// Router Handler / Endpoints

	// users ..
	Router.HandleFunc("/api/user", auth.BasicAuth(registerUserHandler)).Methods("POST")
	Router.HandleFunc("/api/user-login", auth.BasicAuth(postDataLogin)).Methods("POST")
	Router.HandleFunc("/api/user", auth.BasicAuth(getUsersHandler)).Methods("GET")

	// components ..
	Router.Handle("/api/component", auth.JwtMiddleware.Handler(insertCompHandler)).Methods("POST")
	Router.Handle("/api/all-component/{id:[componentID]}", auth.BasicAuth(getAllCompHandler)).Methods("GET")
	Router.Handle("/api/component/{componentID}", auth.BasicAuth(getOneCompHandler)).Methods("GET")
	Router.Handle("/api/component/{componentID}", auth.JwtMiddleware.Handler(updateCompHandler)).Methods("PUT")
	Router.Handle("/api/del-component/{componentID}", auth.JwtMiddleware.Handler(deleteCompHandler)).Methods("PUT")

	//To run the server (same as app.Listen on node.js)
	log.Println(http.ListenAndServe(":8000", Router))

}
