package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Username string
	Token string
}

type UserDetails struct {
	Username string
	Password string
}

type NoteDetails struct {
	Title string
	Description string
}

type DatabaseInterface interface {
	SetupDatabase() error
	Authenticate(username, password string) (string, error)
	GetUserLoginDetails(username string) string
	GetNotes(username string) []NoteDetails
	CreateNotes(token string, newNote NoteDetails) *NoteDetails
}

func NewDatabase() (*DatabaseInterface, error) {
	var db DatabaseInterface = &mockDB{}
	var err error = db.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &db, nil
}