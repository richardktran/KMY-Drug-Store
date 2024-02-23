package routes

import (
	"github.com/gin-gonic/gin"
	todoHandler "github.com/richardktran/MyBlogBE/modules/todo-list/handlers"
)

func V1(router *gin.RouterGroup) {
	todo := router.Group("/todo")
	{
		todo.GET("/:id", todoHandler.GetItemHandler())
	}
}
