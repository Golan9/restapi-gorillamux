package queries

import (
	"encoding/json"
	"net/http"
)

var queries = Queries{}

// UsersQueryHandlers ..
type UsersQueryHandlers struct{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
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
	respondWithJSON(w, http.StatusOK, users)
	return nil
}
