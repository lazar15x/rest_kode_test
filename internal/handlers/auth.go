package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lazar15x/rest_kode_test/api"
	"github.com/lazar15x/rest_kode_test/internal/tools"
)

func (s *Services) Authentication(w http.ResponseWriter, r *http.Request) {
	var credentials tools.UserDetails
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Невозможно декодировать JSON", http.StatusBadRequest)
		api.RequestErrorHandler(w, err)
		return
	}
	defer r.Body.Close()

	authResponse, err := s.db.Authentication(credentials.Username, credentials.Password)
	if err != nil {
		api.RequestErrorHandler(w, err)
		return
	}

	response, _ := json.Marshal(authResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
