package app

import (
	"errors"
	"net/http"
)

type Paging struct {
	Count       int   `json:"count" form:"count"`
	CurrentPage int   `json:"current_page" form:"current_page"`
	PerPage     int   `json:"per_page" form:"per_page"`
	Total       int64 `json:"total" form:"-"`
	TotalPage   int   `json:"total_page" form:"total_page"`
}

type Message struct {
	MessageCode string `json:"message_code"`
	Message     string `json:"message"`
}

type Meta struct {
	Paging Paging `json:"paging,omitempty"`
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

func ResponsePagination(data any, paging Paging) *ResponseType {
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

var (
	ErrorRecordNotFound = errors.New("record not found")
)
