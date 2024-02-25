package handlers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/richardktran/MyBlogBE/app/contracts"
	"github.com/richardktran/MyBlogBE/pkg/app"
)

type TodoHandler struct {
	todoService contracts.TodoService
	userService contracts.UserService
}

func NewTodoHandler(
	todoService contracts.TodoService,
	userService contracts.UserService,
) TodoHandler {
	return TodoHandler{
		todoService: todoService,
		userService: userService,
	}
}

func (h *TodoHandler) GetItemHandler() func(*gin.Context) {
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
