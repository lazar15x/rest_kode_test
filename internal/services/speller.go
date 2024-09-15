package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type SpellCheckResponse struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

var baseURL string = "https://speller.yandex.net/services/spellservice.json/checkText"

func SpellCheck(description string) (string, error) {
	params := url.Values{}
	params.Add("options", "4")
	params.Add("text", description)
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	res, err := http.Get(fullURL)
	if err != nil {
		return description, fmt.Errorf("ошибка отправки запроса: %w", err)
	}
	defer res.Body.Close()

	var spellCheckResults []SpellCheckResponse
	if err := json.NewDecoder(res.Body).Decode(&spellCheckResults); err != nil {
		return description, fmt.Errorf("ошибка декодирвоания ответа: %w", err)
	}

	correctedDescription := applyCorrections(description, spellCheckResults)
	return correctedDescription, nil
}

func applyCorrections(text string, corrections []SpellCheckResponse) string {
	replacementMap := make(map[int]SpellCheckResponse)

	// Создаем карту замен с позицией начала слова
	for _, item := range corrections {
		replacementMap[item.Pos] = item
	}

	var correctedText strings.Builder
	i := 0
	textRune := []rune(text) // Преобразуем текст в руны для правильной работы с юникодом

	// Проходим по тексту
	for i < len(textRune) {
		if correction, found := replacementMap[i]; found {
			// Вставляем исправление
			correctedText.WriteString(correction.S[0])

			// Пропускаем исходное слово (по его длине)
			i += correction.Len
		} else {
			// Добавляем символ как есть
			correctedText.WriteRune(textRune[i])
			i++
		}
	}

	return correctedText.String()
}
