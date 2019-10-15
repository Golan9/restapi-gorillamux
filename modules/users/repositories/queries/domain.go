package queries

import (
	"encoding/json"
	"net/http"
	m "restapi-gorillamux/modules/users/models"
)

var queries = Queries{}
var msg = m.HTTPResponse{}

// UsersQueryHandlers ..
type UsersQueryHandlers struct{}

const (
	empty = ""
	tab   = "\t"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}
func respondWithCode(w http.ResponseWriter, code int, payload interface{}, msg m.HTTPResponse) {
	respondWithJSON(w, code, map[string]interface{}{
		"error":   msg.Error,
		"data":    payload,
		"code":    msg.Code,
		"message": msg.Message,
	})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GetAllUsers ..
func (h *UsersQueryHandlers) GetAllUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := queries.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return err
	}
	msg.Error = err
	msg.Code = http.StatusOK
	msg.Message = "Success"
	respondWithCode(w, http.StatusOK, users, msg)
	return nil
}
