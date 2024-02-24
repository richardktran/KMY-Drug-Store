package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	repositories "github.com/richardktran/MyBlogBE/app/respositories"
	"github.com/richardktran/MyBlogBE/app/services"
	"github.com/richardktran/MyBlogBE/pkg/app"
)

type TodoHandler struct {
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{}
}

func (h *TodoHandler) GetItemHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			app.ResponseBadRequest(err, "invalid_id").Context(c)
			return
		}

		repos := repositories.NewTodoRepository()
		svc := services.NewTodoService(repos)
		data, err := svc.GetItem(id)

		if err != nil {
			app.ResponseNotFound(err, "item_not_found").Context(c)
			return
		}

		app.ResponseSuccess(data).Context(c)
	}
}
