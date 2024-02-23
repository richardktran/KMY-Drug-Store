package app

import (
	"errors"
	"net/http"

	"github.com/richardktran/MyBlogBE/pkg/utils"
)

type Message struct {
	MessageCode string `json:"message_code"`
	Message     string `json:"message"`
}

type Meta struct {
	Paging utils.Paging `json:"paging,omitempty"`
}

type ResponseType struct {
	StatusCode int      `json:"-"`
	Message    *Message `json:"message"`
	Data       any      `json:"data"`
	Meta       *Meta    `json:"meta"`
	RootErr    error    `json:"-"`
	Log        string   `json:"log"`
}

func ResponseSuccess(data any) *ResponseType {
	return &ResponseType{
		StatusCode: http.StatusOK,
		Message: &Message{
			MessageCode: "",
			Message:     "",
		},
		Data:    data,
		Meta:    nil,
		RootErr: nil,
		Log:     "",
	}
}

func ResponsePagination(data any, paging utils.Paging) *ResponseType {
	return &ResponseType{
		StatusCode: http.StatusOK,
		Message: &Message{
			MessageCode: "",
			Message:     "",
		},
		Data: data,
		Meta: &Meta{
			Paging: paging,
		},
		RootErr: nil,
		Log:     "",
	}
}

func ResponseError(statusCode int, root error, message, messageCode string) *ResponseType {
	log := message

	if !IsDebug() {
		root = nil
		log = ""
	}

	if root != nil {
		log = root.Error()
		return &ResponseType{
			StatusCode: statusCode,
			Message:    &Message{MessageCode: messageCode, Message: message},
			RootErr:    root,
			Log:        log,
		}
	}

	return &ResponseType{
		StatusCode: statusCode,
		Message:    &Message{MessageCode: messageCode, Message: message},
		RootErr:    errors.New(message),
		Log:        log,
	}
}

func ResponseBadRequestError(rootError error, message, messageCode string) *ResponseType {

	return ResponseError(
		http.StatusBadRequest,
		rootError,
		message,
		messageCode,
	)
}

func ResponseNotFoundError(rootError error, message, messageCode string) *ResponseType {

	return ResponseError(
		http.StatusNotFound,
		rootError,
		message,
		messageCode,
	)
}

func ResponseInternalServerError(rootError error) *ResponseType {

	return ResponseError(
		http.StatusInternalServerError,
		rootError,
		"Something went wrong",
		"internal_server_error",
	)
}

// ThrowError is a helper function to throw error in services layer

func ThrowError(root error, message, messageCode string, statusCode int) *ResponseType {
	return ResponseError(
		statusCode,
		root,
		message,
		messageCode,
	)
}

func ThrowBadRequestError(root error, message, messageCode string) *ResponseType {
	return ResponseBadRequestError(
		root,
		message,
		messageCode,
	)
}

func ThrowInternalServerError(root error) *ResponseType {
	return ResponseInternalServerError(
		root,
	)
}

func ThrowNotFoundError(root error, message, messageCode string) *ResponseType {
	return ResponseNotFoundError(
		root,
		message,
		messageCode,
	)
}

func (e *ResponseType) RootError() error {
	if err, oke := e.RootErr.(*ResponseType); oke {
		return err.RootError()
	}

	return e.RootErr
}

func (e *ResponseType) Error() string {
	return e.RootErr.Error()
}

var (
	ErrorRecordNotFound = errors.New("record not found")
)
