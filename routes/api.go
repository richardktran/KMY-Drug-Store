package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/richardktran/MyBlogBE/app/controllers"
)

type ApiV1Route struct {
	todoController controllers.TodoController
}

func NewApiV1Route(todoController controllers.TodoController) ApiV1Route {
	return ApiV1Route{
		todoController: todoController,
	}
}

func (r ApiV1Route) Setup(router *gin.Engine) {
	api := router.Group("/api/v1/")
	{
		todo := api.Group("/todo")
		{
			todo.GET("/:id", r.todoController.GetItemController())
		}
	}
}
