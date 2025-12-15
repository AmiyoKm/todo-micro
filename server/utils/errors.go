package utils

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrNotFound             = errors.New("resource not found")
	ErrConflict             = errors.New("resource already exists")
	ErrDuplicateEmail       = errors.New("email already exists")
	ErrDuplicateUsername    = errors.New("username already exists")
	QueryTimeoutDuration    = time.Second * 5
	ErrDuplicatePhoneNumber = errors.New("phone number already exists")
)

type ErrorEnvelop map[string]any

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := ErrorEnvelop{"error": message}

	err := WriteJSON(w, status, env, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "the server encountered a problem and could not process your request"
	ErrorResponse(w, r, 500, message)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	ErrorResponse(w, r, http.StatusNotFound, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	ErrorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func EditConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	ErrorResponse(w, r, http.StatusConflict, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported this resource", r.Method)
	ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func NotPermittedResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user doesn't have the necessary permissions to access this resource"
	ErrorResponse(w, r, http.StatusForbidden, message)
}

func AuthenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this message"
	ErrorResponse(w, r, http.StatusUnauthorized, message)
}
func InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := `invalid authentication credentials`
	ErrorResponse(w, r, http.StatusUnauthorized, message)
}
