package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/richardktran/MyBlogBE/pkg/utils"
)

type Meta struct {
	Paging utils.Paging `json:"paging,omitempty"`
}

type ResponseType struct {
	StatusCode int       `json:"-"`
	Message    *Message  `json:"message"`
	Data       any       `json:"data"`
	Meta       *Meta     `json:"meta"`
	RootErr    *AppError `json:"-"`
	Log        *string   `json:"log,omitempty"`
}

func (e *ResponseType) Context(c *gin.Context) {
	c.JSON(e.StatusCode, e)
}

func ResponseSuccess(data any) *ResponseType {
	return &ResponseType{
		StatusCode: http.StatusOK,
		Data:       data,
	}
}

func ResponsePagination(data any, paging utils.Paging) *ResponseType {
	return &ResponseType{
		StatusCode: http.StatusOK,
		Data:       data,
		Meta: &Meta{
			Paging: paging,
		},
	}
}

func ResponseError(statusCode int, err *AppError) *ResponseType {
	return &ResponseType{
		StatusCode: statusCode,
		Message:    err.Message,
		RootErr:    err,
		Log:        err.Log,
	}
}

func ResponseBadRequest(rootError *AppError) *ResponseType {

	return ResponseError(
		http.StatusBadRequest,
		rootError,
	)
}

func ResponseNotFound(rootError *AppError) *ResponseType {

	return ResponseError(
		http.StatusNotFound,
		rootError,
	)
}

func ResponseInternalServer(rootError *AppError) *ResponseType {

	return ResponseError(
		http.StatusInternalServerError,
		rootError,
	)
}
