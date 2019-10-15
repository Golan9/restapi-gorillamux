package queries

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var queries = Queries{}

// ComponentsQueryHandlers ..
type ComponentsQueryHandlers struct{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GetAllComponent ..
func (h *ComponentsQueryHandlers) GetAllComponent(w http.ResponseWriter, r *http.Request) error {
	params := r.FormValue("componentID")

	if params == "" {
		components, err := queries.FindAll()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return err
		}
		respondWithJSON(w, http.StatusOK, components)
		return nil
	}

	//for i := 0; i < len(params); i++ {
	//componentID := params[i]
	componentID := params
	component, err := queries.FindByID(string(componentID))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return err
	}
	respondWithJSON(w, http.StatusOK, component)

	//}
	return nil
}

// GetOneComponent ..
func (h *ComponentsQueryHandlers) GetOneComponent(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	component, err := queries.FindByID(params["componentID"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return err
	}
	respondWithJSON(w, http.StatusOK, component)
	return nil
}
