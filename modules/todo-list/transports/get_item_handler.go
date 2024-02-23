package transports

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/richardktran/MyBlogBE/pkg/app"
)

func GetItemHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			app.ResponseBadRequest(err, "invalid_id").Context(c)
			return
		}

		app.ResponseSuccess(id).Context(c)
	}
}
