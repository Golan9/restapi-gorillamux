package commands

import (
	"encoding/json"
	"log"
	"net/http"
	database "restapi-gorillamux/database"
	jwt "restapi-gorillamux/jwt"
	m "restapi-gorillamux/modules/users/models"
	q "restapi-gorillamux/modules/users/repositories/queries"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var dao = database.DAO{}
var commands = Commands{}
var queries = q.Queries{}
var msg = m.HTTPResponse{}

// UsersCommandHandlers .. // INI JUGA RAHMAT
type UsersCommandHandlers struct{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithCode(w http.ResponseWriter, code int, payload interface{}, msg m.HTTPResponse) {
	respondWithJSON(w, code, map[string]interface{}{
		"error":   msg.Error,
		"data":    payload,
		"code":    msg.Code,
		"message": msg.Message,
	})
}

// RegisterUser ..
func (h *UsersCommandHandlers) RegisterUser(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	userID := uuid.Must(uuid.NewV4())
	var user m.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request Payload!")

		return err
	}
	user.UserID = userID.String()
	user.Status = m.Active

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

	if err != nil {
		log.Println(err, "Error while generating hash")
		return err
	}
	user.Password = string(hash)

	if err := commands.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return err
	}

	msg.Code = 200
	msg.Error = err
	msg.Message = "Success"

	respondWithCode(w, http.StatusCreated, user, msg)
	return nil
}

// Login ..w http.ResponseWriter, code int, ms
func (h *UsersCommandHandlers) Login(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var user m.User
	var lg m.LoginResponse

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request Payload!")
		return err
	}

	login, err := queries.FindOne(user.Username)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Username or Password")
		return err
	}

	if login.Status == -1 {
		respondWithError(w, http.StatusBadRequest, "User Not Authorized for Login")
		return err
	}

	if login.Status != 1 {
		respondWithError(w, http.StatusBadRequest, "Can't Find User")
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(user.Password))

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Username or Password")
		return err
	}

	token, err := jwt.CreateJwtToken()
	if err != nil {
		log.Println("Error Creating Jwt Token!")
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return err
	}

	lg.Username = login.Username
	lg.Password = login.Password
	lg.Token = token
	msg.Code = 200
	msg.Error = err
	msg.Message = "Success"

	respondWithCode(w, http.StatusOK, lg, msg)
	return nil
}
