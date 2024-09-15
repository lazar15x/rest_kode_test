package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lazar15x/rest_kode_test/api"
	"github.com/lazar15x/rest_kode_test/internal/tools"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			http.Error(w, "Невозможно декодировать JSON", http.StatusBadRequest)
			api.RequestErrorHandler(w, err)
			return
	}
	defer r.Body.Close()

	var db *tools.DatabaseInterface
	var err error
		db, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

	authResponse, err := (*db).Authenticate(credentials.Username, credentials.Password)
	if err != nil {
			// http.Error(w, err.Error(), http.StatusUnauthorized)
			api.RequestErrorHandler(w, err)
			return
	}

	response, _ := json.Marshal(authResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}