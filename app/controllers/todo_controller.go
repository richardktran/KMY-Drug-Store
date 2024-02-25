package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/richardktran/MyBlogBE/app/services/contracts"
	"github.com/richardktran/MyBlogBE/pkg/app"
)

type TodoController struct {
	todoService contracts.ITodoService
	userService contracts.IUserService
}

func NewTodoController(
	todoService contracts.ITodoService,
	userService contracts.IUserService,
) TodoController {
	return TodoController{
		todoService: todoService,
		userService: userService,
	}
}

func (h *TodoController) GetItemController() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			app.ResponseBadRequest(err, "invalid_id").Context(c)
			return
		}

		data, err := h.todoService.GetItem(id)
		user, _ := h.userService.GetUser(1)
		log.Println(user)

		if err != nil {
			app.ResponseNotFound(err, "item_not_found").Context(c)
			return
		}

		app.ResponseSuccess(data).Context(c)
	}
}
