package commands

import (
	"encoding/json"
	"net/http"
	m "restapi-gorillamux/modules/components/models"
	q "restapi-gorillamux/modules/components/repositories/queries"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

var commands = Commands{}
var queries = q.Queries{}

// ComponentsCommandHandlers ..
type ComponentsCommandHandlers struct{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// InsertComponent ..
func (h *ComponentsCommandHandlers) InsertComponent(w http.ResponseWriter, r *http.Request) error {
	// w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	componentID := uuid.Must(uuid.NewV4())
	var component m.Component
	if err := json.NewDecoder(r.Body).Decode(&component); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request Payload!")
		return err
	}
	component.ComponentID = componentID.String()
	component.Status = m.Active
	if err := commands.Insert(component); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return err
	}
	respondWithJSON(w, http.StatusCreated, component)
	return nil
}

// UpdateComponent ..
func (h *ComponentsCommandHandlers) UpdateComponent(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	params := mux.Vars(r)
	component, err := queries.FindByID(params["componentID"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return err
	}
	if err := json.NewDecoder(r.Body).Decode(&component); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return err
	}
	if err := commands.Update(params["componentID"], component); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return err
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	return nil
}

// DeleteComponent ..
func (h *ComponentsCommandHandlers) DeleteComponent(w http.ResponseWriter, r *http.Request) error {
	// w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	component, err := queries.FindByID(params["componentID"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return err
	}
	component.Status = m.Deleted
	if err := commands.Delete(params["componentID"], component); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return err
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	return nil
}
