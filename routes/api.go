package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/richardktran/MyBlogBE/app/handlers"
)

type ApiV1Route struct {
	todoHandler handlers.TodoHandler
}

func NewApiV1Route(todoHandler handlers.TodoHandler) ApiV1Route {
	return ApiV1Route{
		todoHandler: todoHandler,
	}
}

func (r ApiV1Route) Setup(router *gin.Engine) {
	api := router.Group("/api/v1/")
	{
		todo := api.Group("/todo")
		{
			todo.GET("/:id", r.todoHandler.GetItemHandler())
		}
	}
}
