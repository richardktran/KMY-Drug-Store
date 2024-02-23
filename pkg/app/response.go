package app

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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
	Log        *string  `json:"log,omitempty"`
}

func (e *ResponseType) Context(c *gin.Context) {
	c.JSON(e.StatusCode, e)
}

func ResponseSuccess(data any) *ResponseType {
	return &ResponseType{
		StatusCode: http.StatusOK,
		Message: &Message{
			MessageCode: "",
			Message:     "",
		},
		Data: data,
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
	}
}

func ResponseError(statusCode int, root error, messageCode string) *ResponseType {
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

func ResponseBadRequest(rootError error, messageCode string) *ResponseType {

	return ResponseError(
		http.StatusBadRequest,
		rootError,
		messageCode,
	)
}

func ResponseNotFound(rootError error, messageCode string) *ResponseType {

	return ResponseError(
		http.StatusNotFound,
		rootError,
		messageCode,
	)
}

func ResponseInternalServer(rootError error) *ResponseType {

	return ResponseError(
		http.StatusInternalServerError,
		rootError,
		"internal_server_error",
	)
}

// ThrowError is a helper function to throw error in services layer

func ThrowError(root error, messageCode string, statusCode int) *ResponseType {
	return ResponseError(
		statusCode,
		root,
		messageCode,
	)
}

func ThrowBadRequestError(root error, messageCode string) *ResponseType {
	return ResponseBadRequest(
		root,
		messageCode,
	)
}

func ThrowInternalServerError(root error) *ResponseType {
	return ResponseInternalServer(
		root,
	)
}

func ThrowNotFoundError(root error, messageCode string) *ResponseType {
	return ResponseNotFound(
		root,
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
