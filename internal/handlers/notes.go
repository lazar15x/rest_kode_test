package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lazar15x/rest_kode_test/api"
	"github.com/lazar15x/rest_kode_test/internal/services"
	"github.com/lazar15x/rest_kode_test/internal/tools"
	log "github.com/sirupsen/logrus"
)

//Получаем список заметок
func GetNotes(w http.ResponseWriter, r *http.Request) {
	var err error
	var db *tools.DatabaseInterface
	var token = r.Header.Get("Authorization")

	if db, err = tools.NewDatabase(); err != nil {
		api.InternalErrorHandler(w)
		return
	}

	userNotes := (*db).GetNotes(token)
	response := api.NoteResponse{
		Code: http.StatusOK,
		NoteList: userNotes,
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}
}

//Создаем заметку
func CreateNotes(w http.ResponseWriter, r *http.Request) {
	var err error
	var db *tools.DatabaseInterface
	var token = r.Header.Get("Authorization")
	var newNote tools.NoteDetails

	if db, err = tools.NewDatabase(); err != nil {
		api.InternalErrorHandler(w)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil {
		http.Error(w, "Не удалось декодировать JSON", http.StatusBadRequest)
		return
	}
  defer r.Body.Close()
	log.Printf("Received note: %+v", newNote)

	correctedDescription, err := services.SpellCheck(newNote.Description)
	if err != nil {
		log.Printf("Ошибка проверки слов: %v", err)
		return
	}

	newNote.Description = correctedDescription
	addedNote := (*db).CreateNotes(token, newNote)
	
	if addedNote == nil {
		http.Error(w, "Не удалось добавить заметку", http.StatusInternalServerError)
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