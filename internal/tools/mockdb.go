package tools

import (
	"errors"
	"fmt"
)

type mockDB struct{}

var mockUsers = map[string]UserDetails{
	"admin": {
		Username: "admin",
		Password: "admin",
	},
	"user123": {
		Username: "user123",
		Password: "user123",
	},
	"alex": {
		Username: "alex",
		Password: "alex",
	},
	"test": {
		Username: "test",
		Password: "test",
	},
}

var notes = map[string][]NoteDetails{
	"admin": {
		{Title: "Hello", Description: "Hello world"},
		{Title: "Привет мир", Description: "Хаю хай"},
		{Title: "Купить продукты", Description: "молоко, яйца, мясо, творог"},
	},
	"user123": {
		{Title: "Список задач", Description: "1. Погладить рубашку 2. Приготовить поесть"},
	},
	"alex": {
		{Title: "Погулять в 8", Description: "Сходить погуять в 8 часов вечера"},
	},
	"test": {
		{Title: "Тестовое название", Description: "Тестовое описание"},
	},
}

var usersToken = map[string]string{
	"3rte433gggr4":   "admin",
	"fgf5654fgdfg":   "user123",
	"656ffgfgg7676":  "alex",
	"yuuuuui5756756": "test",
}

// Services--------
func (d *mockDB) GetUserLoginDetails(token string) string {
	username, ok := usersToken[token]
	fmt.Println(username)
	if !ok {
		return ""
	}

	return username
}

func (d *mockDB) GetNotes(token string) []NoteDetails {
	username := usersToken[token]
	clientData, exists := notes[username]
	if !exists {
		return nil
	}

	return clientData
}

func (d *mockDB) CreateNotes(token string, newNote NoteDetails) *NoteDetails {
	username := usersToken[token]

	if _, exists := notes[username]; exists {
		notes[username] = append(notes[username], newNote)
	}

	return &newNote
}

func (d *mockDB) Authentication(username, password string) (string, error) {
	var err error
	user, ok := mockUsers[username]
	if !ok || user.Password != password {
		return "", errors.New("неправильный логин или пароль")
	}

	var token string
	for t, u := range usersToken {
		if u == username {
			token = t
			break
		}
	}

	if token == "" {
		return "", errors.New("токен пользователя не найден")
	}

	return token, err
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
