package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/richardktran/MyBlogBE/app/handlers"
)

func V1(router *gin.RouterGroup) {
	todo := router.Group("/todo")
	{
		todo.GET("/:id", handler.NewTodoHandler().GetItemHandler())
	}
}
