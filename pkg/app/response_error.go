package app

import (
	"errors"
	"net/http"

	"github.com/richardktran/KMY-Drug-Store/pkg/utils"
)

type Message struct {
	MessageCode *string `json:"message_code"`
	Message     *string `json:"message"`
}

type AppError struct {
	StatusCode int      `json:"-"`
	Message    *Message `json:"message"`
	RootErr    error    `json:"-"`
	Log        *string  `json:"log,omitempty"`
}

// ThrowError is a helper function to throw error in services layer

func ThrowError(root error, messageCode string, statusCode int) *AppError {
	message := utils.GetMessage(messageCode)
	var log *string

	if !IsDebug() {
		root = nil
		log = nil
	} else {
		log = new(string)
		*log = ""
	}

	if root != nil {
		*log = root.Error()
		return &AppError{
			StatusCode: statusCode,
			Message:    &Message{MessageCode: &messageCode, Message: &message},
			RootErr:    root,
			Log:        log,
		}
	}

	return &AppError{
		StatusCode: statusCode,
		Message:    &Message{MessageCode: &messageCode, Message: &message},
		RootErr:    errors.New(message),
		Log:        log,
	}
}

func ThrowBadRequestError(root error, messageCode string) *AppError {
	return ThrowError(
		root,
		messageCode,
		http.StatusBadRequest,
	)
}

func ThrowInternalServerError(root error) *AppError {
	return ThrowError(
		root,
		"internal_server_error",
		http.StatusInternalServerError,
	)
}

func ThrowNotFoundError(root error, messageCode string) *AppError {
	return ThrowError(
		root,
		messageCode,
		http.StatusNotFound,
	)
}

func ThrowDefaultBadRequestError(root error) *AppError {
	return ThrowError(
		root,
		"invalid_request",
		http.StatusBadRequest,
	)
}

func ThrowDefaultNotFoundError(root error) *AppError {
	return ThrowError(
		root,
		"record_not_found",
		http.StatusNotFound,
	)
}

func (e *AppError) RootError() error {
	if err, oke := e.RootErr.(*AppError); oke {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}
