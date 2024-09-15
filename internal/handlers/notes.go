package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lazar15x/rest_kode_test/api"
	"github.com/lazar15x/rest_kode_test/internal/services"
	"github.com/lazar15x/rest_kode_test/internal/tools"
	log "github.com/sirupsen/logrus"
)

// Получаем список заметок
func (s *Services) GetNotes(w http.ResponseWriter, r *http.Request) {
	var err error
	var token = r.Header.Get("Authorization")

	userNotes := s.db.GetNotes(token)
	response := api.NoteResponse{
		Code:     http.StatusOK,
		NoteList: userNotes,
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}
}

// Создаем заметку
func (s *Services) CreateNotes(w http.ResponseWriter, r *http.Request) {
	var token = r.Header.Get("Authorization")
	var newNote tools.NoteDetails

	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		api.RequestErrorHandler(w, err)
		return
	}
	defer r.Body.Close()

	log.Printf("Получаем заметку: %+v", newNote)
	correctedDescription, err := services.SpellCheck(newNote.Description)

	if err != nil {
		log.Printf("Ошибка проверки слов: %v", err)
		return
	}

	newNote.Description = correctedDescription
	addedNote := s.db.CreateNotes(token, newNote)

	if addedNote == nil {
		api.InternalErrorHandler(w)
		return
	}

	response := api.NoteResponse{
		Code: http.StatusCreated,
		Note: addedNote,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}
}
