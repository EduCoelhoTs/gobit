package validator

import (
	"context"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Validator interface {
	Valid(context.Context) Evaluator
}

type Evaluator map[string]string

// adiciona mensagem de erro no mapa
func (e Evaluator) AddFielError(key, message string) {
	//verifica se o mapa nao existe. se nao existir => cria
	if e == nil {
		e = make(Evaluator)
	}

	//verifica se a chave ja existe no mapa, caso nao exista, adiciona a chave com a mensagem
	if _, ok := e[key]; !ok {
		e[key] = message
	}
}

// AddFielError adds an error message to the Evaluator map for the specified key, if one does not already exist.
// If the map is nil, it initializes the map before adding the error.
// The function does nothing if the key is already present.
// Example usage: e.AddFielError("email", "Invalid email format")
//
// key: the field name for which the error occurred
// message: the error message to associate with the field
func (e *Evaluator) CheckField(ok bool, key, message string) {
	if !ok {
		e.AddFielError(key, message)
	}
}

// NotBlank returns true if value is not a empty string
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars returns true if the number of caracters in value is equal or less than n
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// MinChars returns true if the number of caracters in value is equal or bigger than n
func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

// Matches returns true if the string value matches the provided regular expression pattern.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
