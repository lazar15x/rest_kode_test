package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Username string
	Token    string
}

type UserDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NoteDetails struct {
	Title       string
	Description string
}

type DatabaseInterface interface {
	SetupDatabase() error
	Authentication(username, password string) (string, error)
	GetUserLoginDetails(username string) string
	GetNotes(username string) []NoteDetails
	CreateNotes(token string, newNote NoteDetails) *NoteDetails
}

func NewDatabase() (DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}
	if err := database.SetupDatabase(); err != nil {
		log.Error(err)
		return nil, err
	}

	return database, nil
}
