package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lazar15x/rest_kode_test/internal/tools"
)

type AuthResponse struct {
	Token string
}

type NoteResponse struct {
	Code     int                 `json:"code"`
	NoteList []tools.NoteDetails `json:"noteList,omitempty"`
	Note     *tools.NoteDetails  `json:"note,omitempty"`
}

type Error struct {
	Code    int
	Message string
}

var ErrUnauthorized = errors.New("доступ запрещен")

func writeError(w http.ResponseWriter, message string, code int) {
	res := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Произошла непредвиденная ошибка", http.StatusInternalServerError)
	}
)
